package corbeille

import (
	"encoding/json"
	"net/http"
	corbeillemariadb "portefolio/mariadb/corbeilleMariadb"
	"portefolio/models"
)

func HandleCorbeilleList(w http.ResponseWriter, r *http.Request) {
	if models.SetupCORS(w, r) {
		return
	}

	projets, err := corbeillemariadb.GetCorbeille()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projets)
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

	err := corbeillemariadb.MoveToCorbeille(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}

func HandleGetCorbeilleTech(w http.ResponseWriter, r *http.Request) {
	if models.SetupCORS(w, r) {
		return
	}
	entries, err := corbeillemariadb.GetCorbeilleTech()
	if err != nil {
		http.Error(w, "Erreur récupération corbeille", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(entries)
}
