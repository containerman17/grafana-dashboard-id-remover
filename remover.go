package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	dashboardPath := "/etc/grafana/provisioning/dashboards"
	if len(os.Args) > 1 {
		dashboardPath = os.Args[1]
	} else {
		fmt.Println("No path given, using default path:", dashboardPath)
	}

	files, err := os.ReadDir(dashboardPath)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Starting to remove IDs from dashboard files...")
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".json" {
			filePath := filepath.Join(dashboardPath, file.Name())
			fmt.Printf("Processing %s to remove IDs...\n", file.Name())

			data, err := os.ReadFile(filePath)
			if err != nil {
				fmt.Printf("Error reading %s: %v\n", file.Name(), err)
				continue
			}

			cleanedJSON, err := cleanDashboardJson(data)
			if err != nil {
				fmt.Printf("Error processing %s: %v\n", file.Name(), err)
				continue
			}

			if err := os.WriteFile(filePath, []byte(cleanedJSON), 0644); err != nil {
				fmt.Printf("Error writing %s: %v\n", file.Name(), err)
			}
		}
	}
	fmt.Println("Finished removing IDs from all dashboard files")
}

func cleanDashboardJson(fileContents []byte) ([]byte, error) {
	var dashboard map[string]interface{}
	if err := json.Unmarshal(fileContents, &dashboard); err != nil {
		return nil, fmt.Errorf("error parsing JSON: %v", err)
	}

	delete(dashboard, "uid")
	delete(dashboard, "id")

	if panels, ok := dashboard["panels"].([]interface{}); ok {
		for i := range panels {
			if panel, ok := panels[i].(map[string]interface{}); ok {
				if targets, ok := panel["targets"].([]interface{}); ok {
					for j := range targets {
						if target, ok := targets[j].(map[string]interface{}); ok {
							delete(target, "datasource")
						}
					}
				}
			}
		}
	}

	cleanData, err := json.MarshalIndent(dashboard, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("error converting to JSON: %v", err)
	}

	return cleanData, nil
}
