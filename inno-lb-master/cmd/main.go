package main

import (
	"flag"
	"fmt"
	"inno-lb-master/internal/loadbalancer"
	"inno-lb-master/pkg/config"
	"log"
	"net/http"
)

func main() {

	cfg := config.LoadConfig()
	lb := loadbalancer.NewLoadBalancer()

	configCmd := flag.Bool("config", false, "Print the current configuration")
	statusCmd := flag.Bool("status", false, "Print the current sidecars statuses")
	reloadCmd := flag.Bool("reload", false, "Hot reload configuration")
	flag.Parse()

	if *configCmd {
		lb.PrintConfig()
		return
	}
	if *statusCmd {
		lb.PrintSidecarsStatus()
		return
	}
	if *reloadCmd {
		lb.HotReload()
		return
	}

	http.Handle("/", lb)
	http.HandleFunc("/info", lb.InfoHandler)
	http.HandleFunc("/register", lb.RegisterSidecar)
	http.HandleFunc("/ping", lb.PingSidecar)
	http.HandleFunc("/sidecars", lb.SidecarStatusHandler)
	http.HandleFunc("/reload", lb.ReloadHandler)
	http.HandleFunc("/config", lb.ConfigHandler)

	go lb.MonitorSidecars()

	inputPort := fmt.Sprintf("%d", cfg.PORT)
	fmt.Printf("Load balancer running on port %s\n", inputPort)
	log.Fatal(http.ListenAndServe(":"+inputPort, nil))
}
