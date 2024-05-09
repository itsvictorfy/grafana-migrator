package migrator

import (
	"encoding/json"
	"fmt"
	"log"
)

type Folder struct {
	ID    int    `json:"id"`
	UID   string `json:"uid"`
	Title string `json:"title"`
}

func GetAllFolders(source Grafana) ([]Folder, error) {
	var folders []Folder
	sourceUrl := fmt.Sprintf("%s/api/folders", source)
	body, err := GetReq(sourceUrl, source.ApiToken)
	if err != nil {
		log.Printf("error getting folders %s ", err)
		return nil, fmt.Errorf("error getting folders: %s", err)
	}
	if err := json.Unmarshal([]byte(body), &folders); err != nil {
		log.Println("Error unmarshalling folders:", err)
		return nil, fmt.Errorf("error Unmarshalling dashboard list: %s", err)
	}
	return folders, nil
}

func SetFolders(dest string, apiToken string, folders []Folder) {
	destUrl := fmt.Sprintf("%s/api/folders", dest)
	for _, folder := range folders {
		jsonData, err := json.Marshal(folder)
		if err != nil {
			log.Println("Error:", err)
			return
		}
		body, err := PostReq(destUrl, apiToken, jsonData)
		if err != nil {
			log.Printf("error setting folder %s ", err)
		}
		log.Println(body)
	}
}
