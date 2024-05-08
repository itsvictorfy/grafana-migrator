package migrator

import (
	"encoding/json"
	"fmt"
	"log"
)

type Dashboardinfo struct {
	ID          int      `json:"id"`
	UID         string   `json:"uid"`
	Title       string   `json:"title"`
	URI         string   `json:"uri"`
	URL         string   `json:"url"`
	Slug        string   `json:"slug"`
	Type        string   `json:"type"`
	Tags        []string `json:"tags"`
	IsStarred   bool     `json:"isStarred"`
	SortMeta    int      `json:"sortMeta"`
	FolderID    int      `json:"folderId,omitempty"`
	FolderUID   string   `json:"folderUid,omitempty"`
	FolderTitle string   `json:"folderTitle,omitempty"`
	FolderURL   string   `json:"folderUrl,omitempty"`
}

// func GetDashboardList(body []byte) []DashboardSummary {
// 	var dashboards []DashboardSummary
// 	err := json.Unmarshal(body, &dashboards)
// 	if err != nil {
// 		fmt.Println("Error decoding JSON:", err)
// 		return nil
// 	}
// 	return dashboards
// }

// func ExportDashboards(apiToken, grafanaURL string) {
// 	sourceUrl := fmt.Sprintf("%s/api/search?query=&starred=false", grafanaURL)
// 	err := os.MkdirAll("Dashboards", 0755)
// 	if err != nil {
// 		fmt.Println("Error creating directory:", err)
// 		return
// 	}
// 	body, err := GetReq(sourceUrl, apiToken)
// 	if err != nil {
// 		fmt.Println("Error fetching dashboard:", err)
// 		return
// 	}
// 	dashboards := GetDashboardList(body)
// 	SaveDashboards(apiToken, grafanaURL, dashboards)
// }

//	func SaveDashboards(apiToken, grafanaURL string, dashboards []DashboardSummary) {
//		for _, dashboard := range dashboards {
//			fmt.Printf("ID: %d, UID: %s, Title: %s\n", dashboard.ID, dashboard.UID, dashboard.Title)
//			dashboardURL := fmt.Sprintf("%s/api/dashboards/uid/%s", grafanaURL, dashboard.UID)
//			body, err := GetReq(dashboardURL, apiToken)
//			if err != nil {
//				fmt.Println("Error fetching dashboard:", err)
//				return
//			}
//			filename := fmt.Sprintf("Dashboards/dashboard_%s.json", dashboard.UID)
//			err = os.WriteFile(filename, body, 0644)
//			if err != nil {
//				fmt.Println("Error writing file:", err)
//				return
//			}
//			fmt.Println("Dashboard saved as:", filename)
//		}
//	}
//
// Get a list of all dashbaords -> Source and Apitoken -> retrn Dashboard list, error
func GetDashboardList(source, apiToken string) ([]Dashboardinfo, error) {
	var dashboards []Dashboardinfo
	sourceUrl := fmt.Sprintf("%s/api/search?query=&starred=false", source)
	body, err := GetReq(sourceUrl, apiToken)
	if err != nil {
		log.Println("Error fetching dashboard:", err)
		return nil, fmt.Errorf("error fetching dashboard list: %s", err)
	}
	if err := json.Unmarshal([]byte(body), &dashboards); err != nil {
		log.Println("Error:", err)
		return nil, fmt.Errorf("error Unmarshalling dashboard list: %s", err)
	}
	return dashboards, nil
}

// Get a dashboard by UID -> Source and Apitoken and Dashboard UID -> return Dashboard, error
func GetDashboardByUid(source, apiToken, dbUid string) ([]byte, error) {
	sourceUrl := fmt.Sprintf("%s/api/dashboards/uid/%s", source, dbUid)
	body, err := GetReq(sourceUrl, apiToken)
	if err != nil {
		log.Println("Error fetching dashboard:", err)
		return nil, fmt.Errorf("Error fetching dashboard: %s", err)
	}
	// var jsonBody map[string]interface{}
	// err = json.Unmarshal(body, &jsonBody)
	// if err != nil {
	// 	log.Println("Error parsing JSON: %s", err)
	// 	return nil, fmt.Errorf("Error parsing JSON: %s", err)
	// }
	return body, nil
}
func SetDashboards(dest string, apiToken string, dashboards []Dashboardinfo) error {
	for _, dashboard := range dashboards {
		fmt.Printf("ID: %d, UID: %s, Title: %s\n", dashboard.ID, dashboard.UID, dashboard.Title)
		if dashboard.FolderID == 0 {
			//Dashboard in root folder
			dashboard, err := GetDashboardByUid(dest, apiToken, dashboard.UID)
			if err != nil {
				log.Println("Error fetching dashboard:", err)
				return fmt.Errorf("Error fetching dashboard: %s", err)
			}

			destUrl := fmt.Sprintf("%s/api/dashboards/db", dest)
			PostReq(destUrl, apiToken, dashboard)
		} else {
			//Dashboard in subfolder
		}
	}
	return nil
}
