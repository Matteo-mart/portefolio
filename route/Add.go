package route

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"portefolio/mariadb"
	"portefolio/models"
)

func HandleAddProject(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(50 << 20); err != nil {
		http.Error(w, "Erreur lors de la lecture du formulaire", http.StatusBadRequest)
		return
	}

	title := r.FormValue("title")
	date := r.FormValue("date")
	desc := r.FormValue("description")
	tech := r.FormValue("technologie")
	expl := r.FormValue("explication")
	prob := r.FormValue("probleme")
	sol := r.FormValue("solution")
	urlSource := r.FormValue("url_source")

	result, err := mariadb.DB.Exec(`INSERT INTO project 
        (titre, date_creation, description, technologie, explication, probleme, solution, url_source) 
        VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		title, date, desc, tech, expl, prob, sol, urlSource)
	if err != nil {
		http.Error(w, "Erreur SQL (Project) : "+err.Error(), http.StatusInternalServerError)
		return
	}

	projectID, _ := result.LastInsertId()
	uploadPath := "./templates/uploads/"
	os.MkdirAll(uploadPath, os.ModePerm)

	files := r.MultipartForm.File["image"]
	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			continue
		}
		defer file.Close()

		filename := fileHeader.Filename
		dst, err := os.Create(filepath.Join(uploadPath, filename))
		if err != nil {
			continue
		}
		defer dst.Close()
		io.Copy(dst, file)

		mariadb.DB.Exec(`INSERT INTO project_image (project_id, url) VALUES (?, ?)`, projectID, filename)
	}

	w.WriteHeader(http.StatusCreated)
}

func HandleAddTechnologie(w http.ResponseWriter, r *http.Request) {
	if models.SetupCORS(w, r) {
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	var t models.Technologie
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, "Données invalides", http.StatusBadRequest)
		return
	}

	if t.Nom == "" {
		http.Error(w, "Le nom est requis", http.StatusBadRequest)
		return
	}

	_, err := mariadb.DB.Exec(
		`INSERT INTO technologies (nom, icone, url_source) VALUES (?, ?, ?)`,
		t.Nom, t.Icone, t.Url_source,
	)
	if err != nil {
		log.Printf("Erreur SQL ajout technologie : %v", err)
		http.Error(w, "Erreur SQL : "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Technologie ajoutée"})
}
