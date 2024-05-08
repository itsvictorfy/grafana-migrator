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
	_, err := PostReq(url, apiToken, alerts)
	if err != nil {
		return err
	}

	return nil
}
