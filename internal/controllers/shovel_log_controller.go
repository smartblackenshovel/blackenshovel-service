package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"blackenshovel-service/internal/models"
	"blackenshovel-service/internal/repository"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func ListShovelLogs(w http.ResponseWriter, r *http.Request) {
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))

	var shovelID *uuid.UUID
	if sid := r.URL.Query().Get("shovel_id"); sid != "" {
		if parsed, err := uuid.Parse(sid); err == nil {
			shovelID = &parsed
		}
	}

	var status *string
	if s := r.URL.Query().Get("status"); s != "" {
		status = &s
	}

	logs, err := repository.ListShovelLogs(shovelID, status, limit, offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(logs)
}

func CreateShovelLog(w http.ResponseWriter, r *http.Request) {
	var log models.ShovelLog
	if err := json.NewDecoder(r.Body).Decode(&log); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	log.ID = uuid.New()

	if err := repository.CreateShovelLog(&log); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(log)
}

func GetShovelLogByID(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		http.Error(w, "Invalid UUID", http.StatusBadRequest)
		return
	}

	log, err := repository.GetShovelLogByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if log == nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(log)
}

func DeleteShovelLog(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		http.Error(w, "Invalid UUID", http.StatusBadRequest)
		return
	}

	if err := repository.DeleteShovelLog(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
