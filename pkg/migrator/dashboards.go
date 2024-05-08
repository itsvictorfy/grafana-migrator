package migrator

import (
	"encoding/json"
	"fmt"
	"os"
)

type DashboardSummary struct {
	ID    int    `json:"id"`
	UID   string `json:"uid"`
	Title string `json:"title"`
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

func exportDashboards(apiToken, grafanaURL string) {
	sourceUrl := fmt.Sprintf("%s/api/search?query=&starred=false", grafanaURL)
	err := os.MkdirAll("Dashboards", 0755)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}
	body, err := httpReq("GET", sourceUrl, apiToken, Response{})
	if err != nil {
		fmt.Println("Error fetching dashboard:", err)
		return
	}
	dashboards := getDashboardList(body)
	saveDashboards(apiToken, grafanaURL, dashboards)
}

func saveDashboards(apiToken, grafanaURL string, dashboards []DashboardSummary) {
	for _, dashboard := range dashboards {
		fmt.Printf("ID: %d, UID: %s, Title: %s\n", dashboard.ID, dashboard.UID, dashboard.Title)
		dashboardURL := fmt.Sprintf("%s/api/dashboards/uid/%s", grafanaURL, dashboard.UID)
		body, err := httpReq("GET", dashboardURL, apiToken, Response{})
		if err != nil {
			fmt.Println("Error fetching dashboard:", err)
			return
		}
		filename := fmt.Sprintf("Dashboards/dashboard_%s.json", dashboard.UID)
		err = os.WriteFile(filename, body, 0644)
		if err != nil {
			fmt.Println("Error writing file:", err)
			return
		}
		fmt.Println("Dashboard saved as:", filename)
	}
}
func importDashboards(dashboards []DashboardSummary) {}
