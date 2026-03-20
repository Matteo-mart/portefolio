package route

import (
	"net/http"

	"portefolio/route/corbeille"

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
	r.HandleFunc("/move-to-corbeille", corbeille.HandleMoveToCorbeille).Methods("POST")
	r.HandleFunc("/corbeille-list", corbeille.HandleCorbeilleList).Methods("GET")
	r.HandleFunc("/corbeille-delete", corbeille.HandleCorbeilleDelete).Methods("DELETE")
	r.HandleFunc("/corbeille-restore", corbeille.HandleCorbeilleRestore).Methods("POST")
	r.HandleFunc("/corbeille-vider", corbeille.HandleCorbeilleVider).Methods("DELETE")
	r.HandleFunc("/add-post", HandleAddProject).Methods("POST")
	r.HandleFunc("/delete-project", corbeille.HandleDeleteProject).Methods("DELETE")
	r.HandleFunc("/update-contact", HandleUpdateContact).Methods("PUT", "OPTIONS")
	r.HandleFunc("/update-projet/{id}", HandleUpdateProjet).Methods("PUT", "OPTIONS")
	r.HandleFunc("/update-technologies/{id}", HandleUpdateTechnologies).Methods("PUT", "OPTIONS")
	r.HandleFunc("/delete-technologie", corbeille.HandleDeleteTechnologie).Methods("DELETE", "OPTIONS")
	r.HandleFunc("/corbeille-tech", corbeille.HandleGetCorbeilleTech).Methods("GET")
	r.HandleFunc("/restore-tech", corbeille.HandleRestoreCorbeilleTech).Methods("POST")
	r.HandleFunc("/delete-tech-definitif", corbeille.HandleDeleteDefinitiveTech).Methods("DELETE", "OPTIONS")
	r.HandleFunc("/add-technologie", HandleAddTechnologie).Methods("POST", "OPTIONS")

	r.PathPrefix("/templates/").Handler(
		http.StripPrefix("/templates/", http.FileServer(http.Dir("./templates"))),
	)
}
