package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var apiUrl string = "http://example.com/api"

func main() {
	// define an inputo in  JSON FORMAT
	inputData := `{"name": "John", "age": 30}`

	// validate input JSON
	validatedData, err := validateInput([]byte(inputData))
	if err != nil {
		log.Fatalf("Error validating input: %v", err)
	}

	// record input json
	fmt.Println("Validated input:", validatedData)

	// send HTTP to the API with the valid input
	if err := sendRequest(validatedData); err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
}

// validateInput validate input
func validateInput(input []byte) (map[string]interface{}, error) {
	var data map[string]interface{}
	if err := json.Unmarshal(input, &data); err != nil {
		return nil, err
	}
	// we could add more logic for validation
	return data, nil
}

// function to send the http call
func sendRequest(data map[string]interface{}) error {
	// let's suppose that we will send a call to the API http://example.com/api

	// convert to json format the input
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// send the request to the api
	resp, err := http.Post(apiUrl, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// check the response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API returned non-200 status code: %d", resp.StatusCode)
	}

	fmt.Println("Request sent successfully")
	return nil
}
