package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func PrettyPrintedJSON(data interface{}) string {
	indentedJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatal("Failed to generate JSON: ", err)
	}
	return string(indentedJSON)
}

func SetHeaders(req *http.Request, token string) {
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/vnd.github.v3+json")
}

func HandleResponse(resp *http.Response, obj interface{}) error {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API returned status code %d: %s", resp.StatusCode, body)
	}

	if err := json.Unmarshal(body, &obj); err != nil {
		return err
	}

	return nil
}
