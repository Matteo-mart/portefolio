package route

import (
	"encoding/json"
	"net/http"
	"portefolio/mariadb"
	"portefolio/models"
)

func HandleCorbeilleList(w http.ResponseWriter, r *http.Request) {
	if models.SetupCORS(w, r) {
		return
	}

	projets, err := mariadb.GetCorbeille()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projets)
}

func HandleCorbeilleDelete(w http.ResponseWriter, r *http.Request) {
	if models.SetupCORS(w, r) {
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Paramètre 'id' manquant", http.StatusBadRequest)
		return
	}

	err := mariadb.DeleteFromCorbeille(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}

func HandleCorbeilleRestore(w http.ResponseWriter, r *http.Request) {
	if models.SetupCORS(w, r) {
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Paramètre 'id' manquant", http.StatusBadRequest)
		return
	}

	err := mariadb.RestoreFromCorbeille(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}

func HandleCorbeilleVider(w http.ResponseWriter, r *http.Request) {
	if models.SetupCORS(w, r) {
		return
	}

	err := mariadb.ViderCorbeille()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}

func HandleMoveToCorbeille(w http.ResponseWriter, r *http.Request) {
	if models.SetupCORS(w, r) {
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Paramètre 'id' manquant", http.StatusBadRequest)
		return
	}

	err := mariadb.MoveToCorbeille(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}
