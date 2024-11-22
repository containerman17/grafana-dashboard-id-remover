package main

import (
	_ "embed"
	"encoding/json"
	"testing"
)

//go:embed test_assets/raw.json
var rawData []byte

//go:embed test_assets/expected.json
var expectedData []byte

func TestRemove(t *testing.T) {
	cleanedData, err := cleanDashboardJson(rawData)
	if err != nil {
		t.Fatalf("Failed to clean dashboard JSON: %v", err)
	}

	beautifiedCleaned, err := beautifyJson(cleanedData)
	if err != nil {
		t.Fatalf("Failed to beautify cleaned JSON: %v", err)
	}

	beautifiedExpected, err := beautifyJson(expectedData)
	if err != nil {
		t.Fatalf("Failed to beautify expected JSON: %v", err)
	}

	if string(beautifiedCleaned) != string(beautifiedExpected) {
		t.Errorf("Cleaned JSON does not match expected output")
	}
}

func beautifyJson(data []byte) ([]byte, error) {
	var jsonData interface{}
	if err := json.Unmarshal(data, &jsonData); err != nil {
		return nil, err
	}

	prettyData, err := json.MarshalIndent(jsonData, "", "  ")
	if err != nil {
		return nil, err
	}
	return prettyData, nil
}
