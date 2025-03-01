package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"slices"
	"sync"
	"time"
)

type TestService struct {
	httpClient *http.Client
	groups     []string
	mu         sync.RWMutex
}

func NewTestService() *TestService {
	return &TestService{
		httpClient: &http.Client{},
	}
}

func (ts *TestService) FetchInfo() {
	resp, err := ts.httpClient.Get(os.Getenv("SIDECAR_URL") + "/info")
	if err != nil {
		log.Printf("Error fetching info from sidecar: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Sidecar returned non-OK status: %s\n", resp.Status)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v\n", err)
		return
	}

	var groups []string
	if err := json.Unmarshal(body, &groups); err != nil {
		log.Printf("Error unmarshaling group info: %v\n", err)
		return
	}

	ts.mu.Lock()
	ts.groups = groups
	ts.mu.Unlock()

	log.Printf("Fetched groups from sidecar: %v\n", groups)
}

func (ts *TestService) GetMilk() string {
	req, err := http.NewRequest("GET", os.Getenv("SIDECAR_URL")+"/milk", nil)
	if err != nil {
		return ""
	}

	req.Header.Set("Proxy-Group", "milk")

	resp, err := ts.httpClient.Do(req)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	return string(body)
}

func (ts *TestService) GetDrink() string {
	req, err := http.NewRequest("GET", os.Getenv("SIDECAR_URL"), nil)
	if err != nil {
		return ""
	}

	req.Header.Set("Proxy-Group", "drink")

	resp, err := ts.httpClient.Do(req)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	return string(body)
}

func (ts *TestService) GetSyrup() string {
	req, err := http.NewRequest("GET", os.Getenv("SIDECAR_URL"), nil)
	if err != nil {
		return ""
	}

	req.Header.Set("Proxy-Group", "syrup")

	resp, err := ts.httpClient.Do(req)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	return string(body)
}

func (ts *TestService) CoffeeHandler(w http.ResponseWriter, r *http.Request) {

	var milk, drink, syrup, resp string

	if slices.Contains(ts.groups, "drink") {
		drink = ts.GetDrink()
	} else {
		resp = "There no coffee available\n"
		w.Write([]byte(resp))
		return
	}
	if slices.Contains(ts.groups, "milk") {
		milk = fmt.Sprintf("on %s milk", ts.GetMilk())
	}

	if slices.Contains(ts.groups, "syrup") {
		syrup = fmt.Sprintf("with %s syrup", ts.GetSyrup())
	}

	resp = fmt.Sprintf("Here is your \033[32m%s\033[0m \033[33m%s\033[0m \033[36m%s\033[0m\n", drink, milk, syrup)

	w.Write([]byte(resp))
}

func (ts *TestService) Run() {
	go func() {
		for {
			ts.FetchInfo()
			time.Sleep(10 * time.Second)
		}
	}()
}

func main() {

	testService := NewTestService()
	log.Printf("Test service connecting to sidecar at %s\n", os.Getenv("SIDECAR_URL"))

	testService.Run()

	http.HandleFunc("/coffee", testService.CoffeeHandler)

	port := os.Getenv("PORT")
	log.Printf("Test service running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
