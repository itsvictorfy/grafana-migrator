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

func GetDashboardList(body []byte) []DashboardSummary {
	var dashboards []DashboardSummary
	err := json.Unmarshal(body, &dashboards)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return nil
	}
	return dashboards
}

func ExportDashboards(apiToken, grafanaURL string) {
	sourceUrl := fmt.Sprintf("%s/api/search?query=&starred=false", grafanaURL)
	err := os.MkdirAll("Dashboards", 0755)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}
	body, err := GetReq(sourceUrl, apiToken)
	if err != nil {
		fmt.Println("Error fetching dashboard:", err)
		return
	}
	dashboards := GetDashboardList(body)
	SaveDashboards(apiToken, grafanaURL, dashboards)
}

func SaveDashboards(apiToken, grafanaURL string, dashboards []DashboardSummary) {
	for _, dashboard := range dashboards {
		fmt.Printf("ID: %d, UID: %s, Title: %s\n", dashboard.ID, dashboard.UID, dashboard.Title)
		dashboardURL := fmt.Sprintf("%s/api/dashboards/uid/%s", grafanaURL, dashboard.UID)
		body, err := GetReq(dashboardURL, apiToken)
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
func ImportDashboards(dashboards []DashboardSummary) {}
