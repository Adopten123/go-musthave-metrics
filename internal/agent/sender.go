package agent

import (
	"fmt"
	"net/http"
)

func SendMetrics(s *AgentStorage) {
	baseURL := "http://localhost:8080/update"

	client := &http.Client{}

	for name, value := range s.Gauges {
		url := fmt.Sprintf("%s/gauge/%s/%f", baseURL, name, value)
		sendPOST(client, url)
	}

	for name, value := range s.Counters {
		url := fmt.Sprintf("%s/counter/%s/%d", baseURL, name, value)
		sendPOST(client, url)
	}
}

func sendPOST(client *http.Client, url string) {
	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "text/plain")

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()
}
