package migrator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Preforms POST request get url, apikey, data and return response.Body and error
func PostReq(url, apikey string, data Response) ([]byte, error) {
	client := &http.Client{}

	if len(data.Groups) == 0 {
		log.Println("Error POST creating request: Data is empty")
		return nil, fmt.Errorf("error post creating request: data is empty")
	}
	// Convert data to JSON bytes
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println("Error marshalling PostRequest:", err)
		return nil, fmt.Errorf("error marshalling PostRequest: %w", err)
	}

	// Create a new request with the data in the body
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("Error creating request:", err)
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+apikey)
	req.Header.Set("Content-Type", "application/json") // Set content type to JSON

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending request:", err)
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if resp.StatusCode != 200 {
		if err != nil {
			log.Println("Error sending request:", err)
			return nil, fmt.Errorf("error response status: %s", resp.Status)
		}

		return nil, fmt.Errorf("error response status: %s", resp.Status)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("error reading response body: %w", err)
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	return body, nil
}

// Preforms GET request get url, apikey and return response.Body and error
func GetReq(url, apikey string) ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error creating request:", err)
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+apikey)
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending request:", err)
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if resp.StatusCode != 200 {
		log.Println("Error sending request:", err)
		return nil, fmt.Errorf("error sending request Status code = %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	return body, nil
}
