package main

import (
	"time"

	"github.com/Adopten123/go-musthave-metrics/internal/agent"
)

func main() {
	storage := agent.NewAgentStorage()

	pollTicker := time.NewTicker(2 * time.Second)
	reportTicker := time.NewTicker(10 * time.Second)

	for {
		select {
		case <-pollTicker.C:
			storage.CollectMetrics()
		case <-reportTicker.C:
			agent.SendMetrics(storage)
		}
	}
}
