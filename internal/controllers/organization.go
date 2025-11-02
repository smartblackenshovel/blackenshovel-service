package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"blackenshovel-service/internal/models"
)

var mockOrganizations = []models.Organization{
	{ID: "1", Name: "Acme AG", LegalForm: "AG", Country: "CH"},
	{ID: "2", Name: "Beta GmbH", LegalForm: "GmbH", Country: "CH"},
}

func OrgnizationsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetOrganizations(w, r)
	case http.MethodPost:
		CreateOrganization(w, r)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func GetOrganizations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mockOrganizations)
}

func CreateOrganization(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var org models.Organization

	err := json.NewDecoder(r.Body).Decode(&org)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate required fields
	if org.Name == "" {
		http.Error(w, "Missing required field: name", http.StatusBadRequest)
		return
	}

	// Generate a simple ID for now (in real app, use UUID)
	org.ID = fmt.Sprintf("%d", len(mockOrganizations)+1)

	// Set timestamps
	org.CreatedAt = time.Now()
	org.UpdatedAt = time.Now()

	// Add to mock slice
	mockOrganizations = append(mockOrganizations, org)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(org)
}
