package migrator

import (
	"encoding/json"
	"fmt"
)

func ExportAlerts(apiToken, grafanaURL string) (Response, error) {
	url := fmt.Sprintf("%s/api/v1/provisioning/alert-rules/export", grafanaURL)
	body, err := GetReq(url, apiToken)
	if err != nil {
		return Response{}, err
	}
	var alerts Response
	err = json.Unmarshal(body, &alerts)
	if err != nil {
		return Response{}, err
	}
	return alerts, nil
}

func ImportAlerts(apiToken, grafanaURL string, alerts Response) error {
	url := fmt.Sprintf("%s/api/v1/provisioning/alert-rules", grafanaURL)
	jsonAlert, err := json.Marshal(alerts)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return fmt.Errorf("error marshaling JSON: %s", err)
	}
	_, err = PostReq(url, apiToken, jsonAlert)
	if err != nil {
		return err
	}

	return nil
}
