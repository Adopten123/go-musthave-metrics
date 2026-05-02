package agent

import (
	"fmt"
	"net/http"
)

func SendMetrics(storage *AgentStorage, serverAddr string) {
	client := &http.Client{}

	for name, value := range storage.Gauges {
		url := fmt.Sprintf("http://%s/update/gauge/%s/%f", serverAddr, name, value)
		sendPOST(client, url)
	}

	for name, value := range storage.Counters {
		url := fmt.Sprintf("http://%s/update/counter/%s/%d", serverAddr, name, value)
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
