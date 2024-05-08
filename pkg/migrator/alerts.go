package migrator

import (
	"encoding/json"
	"fmt"
)

func exportAlerts(apiToken, grafanaURL string) (Response, error) {
	url := fmt.Sprintf("%s/api/v1/provisioning/alert-rules/export", grafanaURL)
	body, err := httpReq("GET", url, apiToken, Response{})
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
func importAlerts(apiToken, grafanaURL string, alerts Response) error {
	url := fmt.Sprintf("%s/api/v1/provisioning/alert-rules", grafanaURL)
	_, err := httpReq("POST", url, apiToken, alerts)
	if err != nil {
		return err
	}

	return nil
}
