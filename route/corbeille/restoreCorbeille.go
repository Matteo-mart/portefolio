package corbeille

import (
	"net/http"
	corbeillemariadb "portefolio/mariadb/corbeilleMariadb"
	"portefolio/models"
	"strconv"
)

func HandleCorbeilleRestore(w http.ResponseWriter, r *http.Request) {
	if models.SetupCORS(w, r) {
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Paramètre 'id' manquant", http.StatusBadRequest)
		return
	}

	err := corbeillemariadb.RestoreFromCorbeille(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}

func HandleRestoreCorbeilleTech(w http.ResponseWriter, r *http.Request) {
	if models.SetupCORS(w, r) {
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}

	if err := corbeillemariadb.RestoreFromCorbeilleTech(id); err != nil {
		http.Error(w, "Erreur restauration", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
