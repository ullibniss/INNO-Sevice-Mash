package sidecar

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"sync"
	"time"
)

type Sidecar struct {
	mu         sync.Mutex
	alias      string
	masterURL  *url.URL
	groupInfo  []string
	httpClient *http.Client
}

func NewSidecar() *Sidecar {
	masterURL, err := url.Parse(os.Getenv("MASTER_URL"))
	if err != nil {
		log.Fatalf("Invalid master URL: %v", err)
	}

	return &Sidecar{
		alias:      os.Getenv("ALIAS"),
		masterURL:  masterURL,
		groupInfo:  []string{},
		httpClient: &http.Client{},
	}
}

func (s *Sidecar) FetchGroupInfo() {
	resp, err := s.httpClient.Get(s.masterURL.String() + "/info")
	if err != nil {
		log.Printf("Failed to fetch group info from master: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Master returned non-OK status: %s", resp.Status)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Failed to read response body: %v", err)
		return
	}

	var groups []string
	if err := json.Unmarshal(body, &groups); err != nil {
		log.Printf("Failed to unmarshal group info: %v", err)
		return
	}

	s.mu.Lock()
	s.groupInfo = groups
	s.mu.Unlock()
	log.Println("Group info updated successfully")
}

func (s *Sidecar) ProxyHandler(w http.ResponseWriter, r *http.Request) {
	proxy := httputil.NewSingleHostReverseProxy(s.masterURL)
	r.Host = s.masterURL.Host
	log.Printf("Forwarded request to master: %s", s.masterURL.String())
	proxy.ServeHTTP(w, r)
}

func (s *Sidecar) InfoHandler(w http.ResponseWriter, r *http.Request) {
	s.mu.Lock()
	defer s.mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s.groupInfo)
}

func (s *Sidecar) Register() {

	resp, err := http.Get(s.masterURL.String() + "/register?alias=" + s.alias)
	if err != nil {
		log.Fatalf("Failed to register with master: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Master returned non-OK status: %s", resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(&s.groupInfo); err != nil {
		log.Fatalf("Failed to decode group info: %v", err)
	}

	log.Println("Successfully registered with master and fetched groups")
}

func (s *Sidecar) Ping() {

	for {
		time.Sleep(5 * time.Second) // Ping every 5 seconds
		resp, err := http.Get(s.masterURL.String() + "/ping?alias=" + s.alias)
		if err != nil {
			log.Printf("Failed to send heartbeat to master: %v", err)
			continue
		}
		if err := json.NewDecoder(resp.Body).Decode(&s.groupInfo); err != nil {
			log.Fatalf("Failed to decode group info: %v", err)
		}
		resp.Body.Close()
		log.Println("Heartbeat sent to master")
	}
}
