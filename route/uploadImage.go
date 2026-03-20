package route

import (
	"io"
	"net/http"
	"os"
)

func HandleUpload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(5 << 20)
	file, handler, err := r.FormFile("myImage")
	if err != nil {
		http.Error(w, "Erreur lors de la récupération du fichier", 400)
		return
	}
	defer file.Close()

	filePath := "./templates/uploads/" + handler.Filename
	dst, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Erreur lors de la sauvegarde", 500)
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, "Erreur lors de l écriture", 500)
		return
	}

}
