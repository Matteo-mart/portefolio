package corbeille

import (
	"log"
	"net/http"
	corbeillemariadb "portefolio/mariadb/corbeilleMariadb"
	"portefolio/utils"
	"strconv"
)

func HandleCorbeilleVider(w http.ResponseWriter, r *http.Request) {
	if utils.SetupCORS(w, r) {
		return
	}

	err := corbeillemariadb.ViderCorbeille()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}

func HandleCorbeilleDelete(w http.ResponseWriter, r *http.Request) {
	if utils.SetupCORS(w, r) {
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Paramètre 'id' manquant", http.StatusBadRequest)
		return
	}

	err := corbeillemariadb.DeleteFromCorbeille(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}

func HandleDeleteTechnologie(w http.ResponseWriter, r *http.Request) {
	if utils.SetupCORS(w, r) {
		return
	}
	if r.Method != http.MethodDelete {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}

	if err := corbeillemariadb.MoveToCorbeilleTech(id); err != nil {
		log.Printf("Erreur corbeille tech [%d] : %v", id, err)
		http.Error(w, "Erreur suppression", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func HandleDeleteDefinitiveTech(w http.ResponseWriter, r *http.Request) {
	if utils.SetupCORS(w, r) {
		return
	}
	if r.Method != http.MethodDelete {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}

	if err := corbeillemariadb.DeleteFromCorbeilleTech(id); err != nil {
		http.Error(w, "Erreur suppression définitive", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func HandleDeleteProject(w http.ResponseWriter, r *http.Request) {
	if utils.SetupCORS(w, r) {
		return
	}
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodDelete {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")

	if id == "" {
		http.Error(w, "Erreur: 'id' est vide", http.StatusBadRequest)
		return
	}

	err := corbeillemariadb.DeleteProject(id)
	if err != nil {
		log.Printf("Erreur SQL lors de la suppression de [%s] : %v", id, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
