package service

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"portefolio/internal/domain"
	"strings"
)

type UploadService struct {
	uploadPath  string
	maxFileSize int64
}

// NewUploadService crée un nouveau service d'upload
func NewUploadService(uploadPath string, maxFileSize int64) *UploadService {
	return &UploadService{
		uploadPath:  uploadPath,
		maxFileSize: maxFileSize,
	}
}

// Extensions de fichiers autorisées
var allowedExtensions = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
	".gif":  true,
	".webp": true,
	".pdf":  true,
}

// MIME types autorisés
var allowedMimeTypes = map[string]bool{
	"image/jpeg":      true,
	"image/png":       true,
	"image/gif":       true,
	"image/webp":      true,
	"application/pdf": true,
}

// SaveFile sauvegarde un fichier uploadé de manière sécurisée
func (s *UploadService) SaveFile(fileHeader *multipart.FileHeader) (string, error) {
	// Vérifier la taille
	if fileHeader.Size > s.maxFileSize {
		return "", domain.ErrFileTooLarge
	}

	// Vérifier l'extension
	ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
	if !allowedExtensions[ext] {
		return "", domain.ErrInvalidFileType
	}

	// Ouvrir le fichier
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Lire les premiers bytes pour vérifier le MIME type
	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		return "", err
	}

	// Vérifier le MIME type
	mimeType := fileHeader.Header.Get("Content-Type")
	if !allowedMimeTypes[mimeType] {
		return "", domain.ErrInvalidFileType
	}

	// Reset le reader
	file.Seek(0, 0)

	// Générer un nom de fichier sécurisé (UUID + extension)
	filename, err := s.generateSecureFilename(ext)
	if err != nil {
		return "", err
	}

	// Créer le dossier si nécessaire
	if err := os.MkdirAll(s.uploadPath, os.ModePerm); err != nil {
		return "", err
	}

	// Créer le fichier de destination
	destPath := filepath.Join(s.uploadPath, filename)
	dst, err := os.Create(destPath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	// Copier le fichier
	if _, err := io.Copy(dst, file); err != nil {
		os.Remove(destPath) // Nettoyer en cas d'erreur
		return "", err
	}

	return filename, nil
}

// SaveFiles sauvegarde plusieurs fichiers
func (s *UploadService) SaveFiles(files []*multipart.FileHeader) ([]string, error) {
	var savedFiles []string

	for _, fileHeader := range files {
		filename, err := s.SaveFile(fileHeader)
		if err != nil {
			// En cas d'erreur, supprimer les fichiers déjà sauvegardés
			for _, f := range savedFiles {
				os.Remove(filepath.Join(s.uploadPath, f))
			}
			return nil, err
		}
		savedFiles = append(savedFiles, filename)
	}

	return savedFiles, nil
}

// DeleteFile supprime un fichier
func (s *UploadService) DeleteFile(filename string) error {
	filePath := filepath.Join(s.uploadPath, filename)
	return os.Remove(filePath)
}

// generateSecureFilename génère un nom de fichier sécurisé
func (s *UploadService) generateSecureFilename(ext string) (string, error) {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return fmt.Sprintf("%s%s", hex.EncodeToString(b), ext), nil
}
