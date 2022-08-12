package main

import (
	"testing"
	"strings"
)

func getStatusCode (result string) string {
	topHeader := strings.Split(string(result), "\r\n")[0]

	statusCode := strings.Split(topHeader, " ")[1]

	return statusCode
}

func executeTest (t *testing.T, path string, expectedStatusCode string) {
	result := makeRequest(path)

	statusCode := getStatusCode(result)

	if statusCode != expectedStatusCode {
		t.Errorf("Expected status code %s, got %s", expectedStatusCode, statusCode)
	}
}

func TestValidAsset (t *testing.T) {
	executeTest(t, "/", "200")
}

func TestInvalidAsset (t *testing.T) {
	executeTest(t, "/invalid", "404")
}