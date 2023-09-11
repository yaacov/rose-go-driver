package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/RedHat-Israel/rose-go-driver/pkg/driver"
)

type ResponseData struct {
	Info struct {
		Name   string `json:"name"`
		Action string `json:"action,omitempty"`
	} `json:"info"`
}

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGet(w, r)
	case http.MethodPost:
		handlePost(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	response := ResponseData{
		Info: struct {
			Name   string `json:"name"`
			Action string `json:"action,omitempty"`
		}{
			Name: driver.DriverName,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	var gameData driver.GameData
	err := json.NewDecoder(r.Body).Decode(&gameData)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	action, err := driver.Drive(gameData)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error executing drive method: %v", err), http.StatusInternalServerError)
		return
	}

	response := ResponseData{
		Info: struct {
			Name   string `json:"name"`
			Action string `json:"action,omitempty"`
		}{
			Name:   driver.DriverName,
			Action: action,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
