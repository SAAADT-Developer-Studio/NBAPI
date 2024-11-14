package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type HelloWorldReponse struct {
	Message string `json:"message"`
	Schema  string `json:"$schema"`
}

func TestHandler(t *testing.T) {
	s := &Server{}
	server := httptest.NewServer(http.HandlerFunc(s.healthHandler))
	defer server.Close()
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatalf("error making request to server. Err: %v", err)
	}
	defer resp.Body.Close()
	// Assertions
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", resp.Status)
	}
	// {
	// "$schema": "http://localhost:8080/schemas/GreetingOutputBody.json",
	// "message": "Hello world"
	// }
	data := HelloWorldReponse{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		t.Fatalf("error reading response body. Err: %v", err)
	}
	if data.Message != "Hello world" {
		t.Errorf("expected response body to have message 'Hello world'; got %v", data.Message)
	}
}
