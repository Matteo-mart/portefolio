package domain

import (
	"time"
)

// Project représente un projet du portfolio
type Project struct {
	ID           int64     `json:"id"`
	Titre        string    `json:"title"`
	DateCreation time.Time `json:"date"`
	Description  string    `json:"description"`
	Technologie  string    `json:"technologie"`
	Explication  string    `json:"explication"`
	Probleme     string    `json:"probleme"`
	Solution     string    `json:"solution"`
	UrlSource    string    `json:"url_source"`
	Images       []string  `json:"images,omitempty"`
}

// ProjectImage représente une image associée à un projet
type ProjectImage struct {
	ID        int64  `json:"id"`
	ProjectID int64  `json:"project_id"`
	URL       string `json:"url"`
	MimeType  string `json:"mime_type,omitempty"`
}

// Technology représente une technologie/outil utilisé
type Technology struct {
	ID        int64  `json:"id"`
	Nom       string `json:"nom"`
	Icone     string `json:"icone,omitempty"`
	UrlSource string `json:"url_source,omitempty"`
}

// Contact représente les informations de contact
type Contact struct {
	ID        int64     `json:"id"`
	Telephone string    `json:"telephone,omitempty"`
	Email     string    `json:"email"`
	Linkedin  string    `json:"linkedin,omitempty"`
	Github    string    `json:"github,omitempty"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CorbeilleEntry représente un élément dans la corbeille
type CorbeilleEntry struct {
	ID              int64     `json:"id"`
	ProjectID       int64     `json:"project_id"`
	Titre           string    `json:"titre"`
	DateSuppression time.Time `json:"date_suppression"`
	DateCreation    time.Time `json:"date_creation,omitempty"`
	Description     string    `json:"description,omitempty"`
	Technologie     string    `json:"technologie,omitempty"`
	Explication     string    `json:"explication,omitempty"`
	Probleme        string    `json:"probleme,omitempty"`
	Solution        string    `json:"solution,omitempty"`
	UrlSource       string    `json:"url_source,omitempty"`
	Images          []string  `json:"images,omitempty"`
}

// CorbeilleTech représente une technologie supprimée
type CorbeilleTech struct {
	ID              int64     `json:"id"`
	TechID          int64     `json:"tech_id"`
	Nom             string    `json:"nom"`
	Icone           string    `json:"icone,omitempty"`
	UrlSource       string    `json:"url_source,omitempty"`
	DateSuppression time.Time `json:"date_suppression"`
}

// CreateProjectRequest représente une demande de création de projet
type CreateProjectRequest struct {
	Titre        string `json:"title"`
	DateCreation string `json:"date"`
	Description  string `json:"description"`
	Technologie  string `json:"technologie"`
	Explication  string `json:"explication"`
	Probleme     string `json:"probleme"`
	Solution     string `json:"solution"`
	UrlSource    string `json:"url_source"`
}

// UpdateProjectRequest représente une demande de mise à jour de projet
type UpdateProjectRequest struct {
	Titre        string   `json:"title"`
	Description  string   `json:"description"`
	Technologie  string   `json:"technologie"`
	Explication  string   `json:"explication"`
	Probleme     string   `json:"probleme"`
	Solution     string   `json:"solution"`
	UrlSource    string   `json:"url_source"`
	Images       []string `json:"images,omitempty"`
}

// CreateTechnologyRequest représente une demande de création de technologie
type CreateTechnologyRequest struct {
	Nom       string `json:"nom"`
	Icone     string `json:"icone,omitempty"`
	UrlSource string `json:"url_source,omitempty"`
}

// UpdateTechnologyRequest représente une demande de mise à jour de technologie
type UpdateTechnologyRequest struct {
	Nom       string `json:"nom"`
	Icone     string `json:"icone,omitempty"`
	UrlSource string `json:"url_source,omitempty"`
}

// UpdateContactRequest représente une demande de mise à jour des contacts
type UpdateContactRequest struct {
	ID        int64  `json:"id"`
	Telephone string `json:"telephone,omitempty"`
	Email     string `json:"email"`
	Linkedin  string `json:"linkedin,omitempty"`
	Github    string `json:"github,omitempty"`
}

// Validate valide une demande de création de projet
func (r *CreateProjectRequest) Validate() error {
	if r.Titre == "" {
		return ErrTitreRequired
	}
	if r.DateCreation == "" {
		return ErrDateRequired
	}
	if r.Description == "" {
		return ErrDescriptionRequired
	}
	return nil
}

// Validate valide une demande de création de technologie
func (r *CreateTechnologyRequest) Validate() error {
	if r.Nom == "" {
		return ErrNomRequired
	}
	return nil
}

// Validate valide une demande de mise à jour de contact
func (r *UpdateContactRequest) Validate() error {
	if r.Email == "" {
		return ErrEmailRequired
	}
	return nil
}
