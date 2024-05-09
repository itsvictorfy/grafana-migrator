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

// Get a list of all dashbaords -> Source and Apitoken -> retrn Dashboard list, error
func GetDashboardList(source Grafana) ([]Dashboardinfo, error) {
	var dashboards []Dashboardinfo
	reqUrl := fmt.Sprintf("%s/api/search?query=&starred=false", source.Url)
	body, err := GetReq(reqUrl, source.ApiToken)
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
func GetDashboardByUid(source Grafana, dbUid string) ([]byte, error) {
	reqUrl := fmt.Sprintf("%s/api/dashboards/uid/%s", source, dbUid)
	body, err := GetReq(reqUrl, source.ApiToken)
	if err != nil {
		log.Println("Error fetching dashboard:", err)
		return nil, fmt.Errorf("error fetching dashboard: %s", err)
	}
	// var jsonBody map[string]interface{}
	// err = json.Unmarshal(body, &jsonBody)
	// if err != nil {
	// 	log.Println("Error parsing JSON: %s", err)
	// 	return nil, fmt.Errorf("Error parsing JSON: %s", err)
	// }
	return body, nil
}
func SetDashboard(dest Grafana, dashboard []byte) error {
	destUrl := fmt.Sprintf("%s/api/dashboards/db", dest)
	// jsonDashboard, err := json.Marshal(dashboard)
	// if err != nil {
	// 	fmt.Println("Error marshaling JSON:", err)
	// 	return fmt.Errorf("Error marshaling JSON: %s", err)
	// }
	PostReq(destUrl, dest.ApiToken, dashboard)

	return nil
}
