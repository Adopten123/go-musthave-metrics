package main

import (
	"flag"
	"os"
	"strconv"
	"time"

	"github.com/Adopten123/go-musthave-metrics/internal/agent"
)

func main() {
	var serverAddr string
	var reportInterval int
	var pollInterval int

	flag.StringVar(&serverAddr, "a", "localhost:8080", "address and port of the server")
	flag.IntVar(&reportInterval, "r", 10, "report interval in seconds")
	flag.IntVar(&pollInterval, "p", 2, "poll interval in seconds")
	flag.Parse()

	if envAddr := os.Getenv("ADDRESS"); envAddr != "" {
		serverAddr = envAddr
	}
	if envReport := os.Getenv("REPORT_INTERVAL"); envReport != "" {
		if val, err := strconv.Atoi(envReport); err == nil {
			reportInterval = val
		}
	}
	if envPoll := os.Getenv("POLL_INTERVAL"); envPoll != "" {
		if val, err := strconv.Atoi(envPoll); err == nil {
			pollInterval = val
		}
	}

	pollTicker := time.NewTicker(time.Duration(pollInterval) * time.Second)
	reportTicker := time.NewTicker(time.Duration(reportInterval) * time.Second)

	defer pollTicker.Stop()
	defer reportTicker.Stop()

	storage := agent.NewAgentStorage()

	for {
		select {
		case <-pollTicker.C:
			storage.CollectMetrics()
		case <-reportTicker.C:
			agent.SendMetrics(storage, serverAddr)
		}
	}
}
