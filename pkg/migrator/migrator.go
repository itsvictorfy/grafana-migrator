package migrator

import (
	"fmt"
	"log"
)

func MigrateFolders(source, dest Grafana) {
}
func MigrateDashboards(source, dest Grafana) error {
	dashboards, err := GetDashboardList(source)
	if err != nil {
		log.Printf("Error during dahsboard export: %s", err)
		return fmt.Errorf("error during dahsboard export: %s", err)
	}
	for _, dashboard := range dashboards {
		fmt.Printf("ID: %d, UID: %s, Title: %s\n", dashboard.ID, dashboard.UID, dashboard.Title)
		if dashboard.FolderID == 0 {
			//Dashboard in root folder
			dashboard, err := GetDashboardByUid(source, dashboard.UID)
			if err != nil {
				log.Println("Error fetching dashboard:", err)
				return fmt.Errorf("error fetching dashboard: %s", err)
			}
			err = SetDashboard(dest, dashboard)
			if err != nil {
				log.Printf("Error during dahsboard import %s", err)
				return fmt.Errorf("error during dahsboard import %s", err)
			}
		}
	}
	return nil
}
func MigrateAlers(source, dest Grafana) {}
