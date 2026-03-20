package route

import (
	"net/http"

	"github.com/gorilla/mux"
)

func DefRoute(r *mux.Router) {

	r.HandleFunc("/", HandleHome).Methods("GET")
	r.HandleFunc("/contact.html", HandleContact).Methods("GET")
	r.HandleFunc("/projet.html", HandleProjet).Methods("GET")
	r.HandleFunc("/project/{id}", GetProjectHandler).Methods("GET")
	r.HandleFunc("/admin", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/admin.html")
	}).Methods("GET")
	r.HandleFunc("/corbeille", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/corbeille.html")
	}).Methods("GET")
	r.PathPrefix("/uploads/").Handler(
		http.StripPrefix("/uploads/", http.FileServer(http.Dir("./templates/uploads"))),
	)

	// CORBEILLE
	r.HandleFunc("/move-to-corbeille", HandleMoveToCorbeille).Methods("POST")
	r.HandleFunc("/corbeille-list", HandleCorbeilleList).Methods("GET")
	r.HandleFunc("/corbeille-delete", HandleCorbeilleDelete).Methods("DELETE")
	r.HandleFunc("/corbeille-restore", HandleCorbeilleRestore).Methods("POST")
	r.HandleFunc("/corbeille-vider", HandleCorbeilleVider).Methods("DELETE")
	r.HandleFunc("/add-post", HandleAddProject).Methods("POST")
	r.HandleFunc("/delete-project", HandleDeleteProject).Methods("DELETE")
	r.HandleFunc("/update-contact", HandleUpdateContact).Methods("PUT", "OPTIONS")
	r.HandleFunc("/update-projet/{id}", HandleUpdateProjet).Methods("PUT", "OPTIONS")
	r.HandleFunc("/update-technologies/{id}", HandleUpdateTechnologies).Methods("PUT", "OPTIONS")
	r.HandleFunc("/delete-technologie", HandleDeleteTechnologie).Methods("DELETE", "OPTIONS")
	r.HandleFunc("/corbeille-tech", HandleGetCorbeilleTech).Methods("GET")
	r.HandleFunc("/restore-tech", HandleRestoreCorbeilleTech).Methods("POST")
	r.HandleFunc("/delete-tech-definitif", HandleDeleteDefinitiveTech).Methods("DELETE", "OPTIONS")
	r.HandleFunc("/add-technologie", HandleAddTechnologie).Methods("POST", "OPTIONS")

	r.PathPrefix("/templates/").Handler(
		http.StripPrefix("/templates/", http.FileServer(http.Dir("./templates"))),
	)
}
