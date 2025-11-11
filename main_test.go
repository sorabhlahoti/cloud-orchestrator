package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(healthHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response map[string]string
	json.NewDecoder(rr.Body).Decode(&response)
	if response["status"] != "healthy" {
		t.Errorf("handler returned unexpected body: got %v want %v", response["status"], "healthy")
	}
}

func TestProvisionHandler(t *testing.T) {
	req, err := http.NewRequest("POST", "/provision", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(provisionHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	var resource Resource
	json.NewDecoder(rr.Body).Decode(&resource)
	if resource.ID == 0 {
		t.Error("Expected resource to have an ID")
	}
	if resource.Name == "" {
		t.Error("Expected resource to have a name")
	}
}

func TestListHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/resources", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(listHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var resList []Resource
	json.NewDecoder(rr.Body).Decode(&resList)
	// List can be empty, just check it's a valid array
	if resList == nil {
		t.Error("Expected resources list to be initialized")
	}
}

func TestDeleteHandler(t *testing.T) {
	// First provision a resource
	resources[12345] = Resource{ID: 12345, Name: "test-resource"}

	req, err := http.NewRequest("DELETE", "/resources/12345", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(deleteHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Verify resource was deleted
	if _, exists := resources[12345]; exists {
		t.Error("Resource should have been deleted")
	}
}

func TestDeleteHandlerNotFound(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/resources/999999", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(deleteHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}
}

func TestProvisionHandlerInvalidMethod(t *testing.T) {
	req, err := http.NewRequest("GET", "/provision", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(provisionHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusMethodNotAllowed)
	}
}
