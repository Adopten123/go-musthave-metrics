package agent

import (
	"math/rand"
	"runtime"
)

type AgentStorage struct {
	Gauges   map[string]float64
	Counters map[string]int64
}

func NewAgentStorage() *AgentStorage {
	return &AgentStorage{
		Gauges:   make(map[string]float64),
		Counters: make(map[string]int64),
	}
}

func (s *AgentStorage) CollectMetrics() {
	var rtm runtime.MemStats
	runtime.ReadMemStats(&rtm)

	s.Gauges["Alloc"] = float64(rtm.Alloc)
	s.Gauges["BuckHashSys"] = float64(rtm.BuckHashSys)
	s.Gauges["Frees"] = float64(rtm.Frees)
	s.Gauges["GCCPUFraction"] = rtm.GCCPUFraction
	s.Gauges["GCSys"] = float64(rtm.GCSys)
	s.Gauges["HeapAlloc"] = float64(rtm.HeapAlloc)
	s.Gauges["HeapIdle"] = float64(rtm.HeapIdle)
	s.Gauges["HeapInuse"] = float64(rtm.HeapInuse)
	s.Gauges["HeapObjects"] = float64(rtm.HeapObjects)
	s.Gauges["HeapReleased"] = float64(rtm.HeapReleased)
	s.Gauges["HeapSys"] = float64(rtm.HeapSys)
	s.Gauges["LastGC"] = float64(rtm.LastGC)
	s.Gauges["Lookups"] = float64(rtm.Lookups)
	s.Gauges["MCacheInuse"] = float64(rtm.MCacheInuse)
	s.Gauges["MCacheSys"] = float64(rtm.MCacheSys)
	s.Gauges["MSpanInuse"] = float64(rtm.MSpanInuse)
	s.Gauges["MSpanSys"] = float64(rtm.MSpanSys)
	s.Gauges["Mallocs"] = float64(rtm.Mallocs)
	s.Gauges["NextGC"] = float64(rtm.NextGC)
	s.Gauges["NumForcedGC"] = float64(rtm.NumForcedGC)
	s.Gauges["NumGC"] = float64(rtm.NumGC)
	s.Gauges["OtherSys"] = float64(rtm.OtherSys)
	s.Gauges["PauseTotalNs"] = float64(rtm.PauseTotalNs)
	s.Gauges["StackInuse"] = float64(rtm.StackInuse)
	s.Gauges["StackSys"] = float64(rtm.StackSys)
	s.Gauges["Sys"] = float64(rtm.Sys)
	s.Gauges["TotalAlloc"] = float64(rtm.TotalAlloc)

	s.Gauges["RandomValue"] = rand.Float64()
	s.Counters["PollCount"]++
}
