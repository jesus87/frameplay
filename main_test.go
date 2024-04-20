package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

// simplete test for validation function
func TestValidateInput(t *testing.T) {
	// valid input
	validInput := []byte(`{"name": "John", "age": 30}`)
	_, err := validateInput(validInput)
	if err != nil {
		t.Errorf("Expected no error for valid input, got %v", err)
	}

}

// simple test for  HTTP CALL
func TestSendRequest(t *testing.T) {
	// Mock data for request
	mockData := map[string]interface{}{
		"key": "value",
	}

	// lets create a server for simulation
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// validate the post
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST request, got %s", r.Method)
		}

		// validate the content
		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("Expected Content-Type to be application/json, got %s", r.Header.Get("Content-Type"))
		}

		// read the body
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		body := buf.String()

		expectedBody := `{"key":"value"}`
		if body != expectedBody {
			t.Errorf("Expected request body %s, got %s", expectedBody, body)
		}

		// return with status ok
		w.WriteHeader(http.StatusOK)
	}))

	defer server.Close()

	// change the url for our test server
	apiUrl = server.URL

	// send the http call
	err := sendRequest(mockData)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
