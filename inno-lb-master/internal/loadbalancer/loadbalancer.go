package loadbalancer

import (
	"encoding/json"
	"fmt"
	"inno-lb-master/internal/endpoint"
	"inno-lb-master/internal/group"
	"inno-lb-master/internal/sidecar"
	"inno-lb-master/pkg/config"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
	"time"
)

type LoadBalancer struct {
	mu          sync.Mutex
	groups      []*group.Group
	groupMap    map[string]*group.Group
	totalWeight map[string]int
	counters    map[string]int
	Sidecars    map[string]*sidecar.Sidecar `yaml:"sidecars"`
	currentIdx  map[string]int
}

func NewLoadBalancer() *LoadBalancer {
	cfg := config.LoadConfig()

	lb := &LoadBalancer{
		groupMap:    make(map[string]*group.Group),
		totalWeight: make(map[string]int),
		counters:    make(map[string]int),
		currentIdx:  make(map[string]int),
		Sidecars:    make(map[string]*sidecar.Sidecar),
	}

	for _, group := range cfg.GROUPS {
		for _, endpoint := range group.Endpoints {
			weight := endpoint.Weight
			lb.totalWeight[group.Alias] += weight
		}
		lb.groups = append(lb.groups, group)
		lb.groupMap[group.Alias] = group
		lb.counters[group.Alias] = 0
		lb.currentIdx[group.Alias] = 0
	}

	return lb
}

func (lb *LoadBalancer) getNextEndpoint(alias string) *endpoint.Endpoint {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	group := lb.groupMap[alias]
	for {
		lb.counters[alias]++
		if lb.counters[alias] >= lb.totalWeight[alias] {
			lb.counters[alias] = 0
			lb.currentIdx[alias] = (lb.currentIdx[alias] + 1) % len(group.Endpoints)
		}

		endpoint := group.Endpoints[lb.currentIdx[alias]]
		if lb.counters[alias] < endpoint.Weight {
			return endpoint
		}
	}
}

func (lb *LoadBalancer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	alias := r.Header.Get("Proxy-Group")
	if alias == "" {
		http.Error(w, "Missing Proxy-Group header", http.StatusBadRequest)
		return
	}

	_, exists := lb.groupMap[alias]
	if !exists {
		http.Error(w, fmt.Sprintf("Group not found: %s", alias), http.StatusNotFound)
		return
	}

	endpoint := lb.getNextEndpoint(alias)
	endpointUrl, _ := url.Parse(endpoint.URL)
	proxy := httputil.NewSingleHostReverseProxy(endpointUrl)
	r.Host = endpointUrl.Host
	proxy.ServeHTTP(w, r)

	fmt.Printf("Forwarded request to group: %s, endpoint: %s\n", alias, endpointUrl.String())
}

func (lb *LoadBalancer) InfoHandler(w http.ResponseWriter, r *http.Request) {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	groups := []string{}
	for _, group := range lb.groups {
		groups = append(groups, group.Alias)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(groups)
}

func (lb *LoadBalancer) RegisterSidecar(w http.ResponseWriter, r *http.Request) {
	address := r.RemoteAddr
	alias := r.URL.Query().Get("alias")

	if alias == "" {
		http.Error(w, "Failed to identify sidecar", http.StatusBadRequest)
		return
	}

	lb.mu.Lock()
	defer lb.mu.Unlock()

	if _, exists := lb.Sidecars[alias]; exists {
		lb.Sidecars[alias].Status = "UP"
		log.Printf("Sidecar %s at %s reconnected\n", alias, address)
	} else {
		lb.Sidecars[alias] = &sidecar.Sidecar{
			Address: address,
			Alias:   alias,
			Status:  "UP",
		}
		log.Printf("Registered sidecar %s at %s\n", alias, address)
	}

	groups := []string{}
	for _, group := range lb.groups {
		groups = append(groups, group.Alias)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(groups)

}

func (lb *LoadBalancer) PingSidecar(w http.ResponseWriter, r *http.Request) {
	address := r.RemoteAddr
	alias := r.URL.Query().Get("alias")
	lb.mu.Lock()
	defer lb.mu.Unlock()

	if sidecar, exists := lb.Sidecars[alias]; exists {
		sidecar.LastPing = time.Now()
		sidecar.Status = "UP"
	} else {
		http.Error(w, "Sidecar not registered", http.StatusBadRequest)
		return
	}

	groups := []string{}
	for _, group := range lb.groups {
		groups = append(groups, group.Alias)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(groups)
	log.Printf("Heartbeat received from sidecar %s at %s\n", alias, address)
}

func (lb *LoadBalancer) MonitorSidecars() {
	for range time.Tick(5 * time.Second) {
		lb.mu.Lock()

		for alias, sidecar := range lb.Sidecars {
			if time.Since(sidecar.LastPing) > 10*time.Second {
				sidecar.Status = "DOWN"
				log.Printf("Sidecar at %s is DOWN\n", alias)
			}
		}

		lb.mu.Unlock()
	}
}

func (lb *LoadBalancer) ConfigHandler(w http.ResponseWriter, r *http.Request) {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	cfg := config.LoadConfig()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cfg)
}

func (lb *LoadBalancer) SidecarStatusHandler(w http.ResponseWriter, r *http.Request) {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	statuses := make(map[string]string)
	for alias, sidecar := range lb.Sidecars {
		statuses[alias] = sidecar.Status
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(statuses)
}

func (lb *LoadBalancer) ReloadHandler(w http.ResponseWriter, r *http.Request) {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	cfg := config.ReloadConfig()

	lb.groups = cfg.GROUPS
	for _, group := range cfg.GROUPS {
		for _, endpoint := range group.Endpoints {
			weight := endpoint.Weight
			lb.totalWeight[group.Alias] += weight
		}
		lb.groupMap[group.Alias] = group
		lb.counters[group.Alias] = 0
		lb.currentIdx[group.Alias] = 0
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("OK")
}

func (lb *LoadBalancer) PrintConfig() {
	cfg := config.LoadConfig()

	resp, err := http.Get(fmt.Sprintf("http://127.0.0.1:%d/config", cfg.PORT))
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}

	fmt.Println(string(body))
}

func (lb *LoadBalancer) PrintSidecarsStatus() {

	cfg := config.LoadConfig()
	resp, err := http.Get(fmt.Sprintf("http://127.0.0.1:%d/sidecars", cfg.PORT))
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}

	fmt.Println(string(body))
}

func (lb *LoadBalancer) HotReload() {
	cfg := config.LoadConfig()
	_, err := http.Get(fmt.Sprintf("http://127.0.0.1:%d/reload", cfg.PORT))
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		return
	}
	fmt.Println("Reloaded")
}
