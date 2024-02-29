package migrator

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

var apiToken = "glsa_Y0PX0CL5zEKsl0wUQv3y617kBHaKfHcf_49312333"
var grafanaURL = "https://grafana-qa.50coins.com"

type DashboardSummary struct {
	ID    int    `json:"id"`
	UID   string `json:"uid"`
	Title string `json:"title"`
}

func httpReq(method, url, apikey string) []byte {
	// Create HTTP client
	client := &http.Client{}

	// Create GET request
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return []byte(err.Error())
	}
	req.Header.Set("Authorization", "Bearer "+apikey)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return []byte(err.Error())
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return []byte(err.Error())
	}
	return body
}

func getDashboardList(body []byte) []DashboardSummary {
	var dashboards []DashboardSummary
	err := json.Unmarshal(body, &dashboards)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil
	}
	return dashboards
}
func exportDashboards(dashboards []DashboardSummary) {
	for _, dashboard := range dashboards {
		fmt.Printf("ID: %d, UID: %s, Title: %s\n", dashboard.ID, dashboard.UID, dashboard.Title)
		dashboardURL := fmt.Sprintf("%s/api/dashboards/uid/%s", grafanaURL, dashboard.UID)
		body := httpReq("GET", dashboardURL, apiToken)

		filename := fmt.Sprintf("Dashboards/dashboard_%s.json", dashboard.UID)
		err := os.WriteFile(filename, body, 0644)
		if err != nil {
			fmt.Println("Error writing file:", err)
			return
		}
		fmt.Println("Dashboard saved as:", filename)
	}
}
func getDashboards() {
	sourceUrl := fmt.Sprintf("%s/api/search?query=&starred=false", grafanaURL)
	err := os.MkdirAll("Dashboards", 0755)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}
	body := httpReq("GET", sourceUrl, apiToken)
	dashboards := getDashboardList(body)
	exportDashboards(dashboards)
}
