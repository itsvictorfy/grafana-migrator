package migrator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func httpReq(method, url, apikey string, data Response) ([]byte, error) {
	// Create HTTP client
	client := &http.Client{}
	if method == "POST" {
		if len(data.Groups) == 0 {
			fmt.Println("Error POST creating request: Data is empty")
			return nil, fmt.Errorf("error post creating request: data is empty")
		}
		// Convert data to JSON bytes
		jsonData, err := json.Marshal(data)
		if err != nil {
			fmt.Println("Error marshalling data:", err)
			return nil, err
		}

		// Create a new request with the data in the body
		req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))
		if err != nil {
			fmt.Println("Error creating request:", err)
			return nil, err
		}
		req.Header.Set("Authorization", "Bearer "+apikey)
		req.Header.Set("Content-Type", "application/json") // Set content type to JSON

		// Send the request
		resp, err := client.Do(req)
		if resp.StatusCode != 200 {
			if err != nil {
				fmt.Println("Error sending request:", err)
				return nil, err
			}

			return nil, fmt.Errorf("error sending request Status code = %d", resp.StatusCode)
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return nil, err
		}
		return body, nil
	}
	if method == "GET" {
		req, err := http.NewRequest(method, url, nil)
		if err != nil {
			fmt.Println("Error creating request:", err)
			return nil, err
		}
		req.Header.Set("Authorization", "Bearer "+apikey)
		resp, err := client.Do(req)
		if resp.StatusCode != 200 {
			if err != nil {
				fmt.Println("Error sending request:", err)
				return nil, err
			}

			return nil, fmt.Errorf("error sending request Status code = %d", resp.StatusCode)
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return nil, err
		}
		return body, nil
	}
	return nil, fmt.Errorf("invalid http method: %s", method)
}
