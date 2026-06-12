package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

// Response représente une réponse API standardisée
type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// WriteJSON écrit une réponse JSON
func WriteJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteStatus(status)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Erreur lors de l'encodage JSON: %v", err)
	}
}

// WriteSuccess écrit une réponse JSON de succès
func WriteSuccess(w http.ResponseWriter, status int, data interface{}) {
	WriteJSON(w, status, Response{
		Success: true,
		Data:    data,
	})
}

// WriteError écrit une réponse JSON d'erreur
func WriteError(w http.ResponseWriter, status int, message string) {
	WriteJSON(w, status, Response{
		Success: false,
		Error:   message,
	})
}

// WriteValidationError écrit une erreur de validation
func WriteValidationError(w http.ResponseWriter, message string) {
	WriteError(w, http.StatusBadRequest, message)
}

// WriteInternalError écrit une erreur interne
func WriteInternalError(w http.ResponseWriter, err error) {
	log.Printf("Erreur interne: %v", err)
	WriteError(w, http.StatusInternalServerError, "Une erreur interne s'est produite")
}

// WriteNotFound écrit une erreur 404
func WriteNotFound(w http.ResponseWriter, message string) {
	WriteError(w, http.StatusNotFound, message)
}

// ParseJSON décode un body JSON
func ParseJSON(r *http.Request, v interface{}) error {
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(v)
}
