package domain

import "errors"

// Erreurs de validation
var (
	ErrTitreRequired       = errors.New("le titre est requis")
	ErrDateRequired        = errors.New("la date est requise")
	ErrDescriptionRequired = errors.New("la description est requise")
	ErrNomRequired         = errors.New("le nom est requis")
	ErrEmailRequired       = errors.New("l'email est requis")
)

// Erreurs métier
var (
	ErrProjectNotFound    = errors.New("projet non trouvé")
	ErrTechnologyNotFound = errors.New("technologie non trouvée")
	ErrContactNotFound    = errors.New("contact non trouvé")
	ErrInvalidID          = errors.New("ID invalide")
	ErrInvalidFileType    = errors.New("type de fichier non autorisé")
	ErrFileTooLarge       = errors.New("fichier trop volumineux")
)
