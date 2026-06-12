# Portfolio

Application web de gestion de portfolio personnel développée en Go avec Gorilla Mux et MariaDB.

## Description

Ce projet est une application web permettant de présenter et gérer un portfolio de projets. Il offre une interface d'administration complète pour ajouter, modifier et supprimer des projets, ainsi qu'une gestion des technologies et des contacts.

### Fonctionnalités principales

- Affichage public du portfolio avec projets et technologies
- Interface d'administration pour gérer les projets
- Gestion des contacts (téléphone, email, LinkedIn, GitHub)
- Système de corbeille pour restaurer les projets supprimés
- Upload et gestion d'images pour chaque projet
- Support multimédia avec types MIME

## Architecture

```
portefolio/
├── main.go              # Point d'entrée de l'application
├── go.mod               # Dépendances Go
├── mariadb/             # Couche d'accès aux données
│   ├── connection.go    # Connexion à la base de données
│   ├── All.go           # Requêtes de lecture
│   ├── getProjectByID.go
│   └── update.go        # Requêtes de modification
├── route/               # Handlers HTTP
│   ├── route.go         # Configuration des routes
│   ├── Home.go          # Page d'accueil
│   ├── projet.go        # Gestion des projets
│   ├── Add.go           # Ajout de données
│   └── update.go        # Mise à jour
├── templates/           # Vues HTML
│   ├── home.html
│   ├── admin.html
│   ├── projet.html
│   ├── contact.html
│   └── corbeille.html
└── utils/               # Utilitaires
    ├── struct.go        # Structures de données
    ├── cors.go          # Configuration CORS
    ├── clear.go         # Utilitaires terminal
    └── baseDeDonnée.sql # Schéma de la base de données
```

## Technologies utilisées

- **Backend**: Go 1.25.0
- **Router**: Gorilla Mux
- **Base de données**: MariaDB
- **Driver SQL**: go-sql-driver/mysql

## Prérequis

- Go 1.25.0 ou supérieur
- MariaDB
- Système Linux (OpenSUSE recommandé pour le script d'installation)

## Installation

### 1. Installation de MariaDB

Le projet inclut un script d'installation automatique pour MariaDB :

```bash
chmod +x mariadbIN.sh
sudo ./mariadbIN.sh
```

Ce script va :
- Mettre à jour le système
- Installer MariaDB
- Démarrer le service MariaDB
- Se connecter avec l'utilisateur `matteo`

### 2. Configuration de la base de données

Une fois connecté à MariaDB, exécutez le script SQL :

```bash
mariadb -u matteo -p < utils/baseDeDonnée.sql
```

Ou manuellement dans le shell MariaDB :

```sql
source utils/baseDeDonnée.sql
```

Le schéma créera :
- Base de données `portefolio`
- Table `project` (projets)
- Table `project_image` (images des projets)
- Table `contacts` (informations de contact)
- Table `corbeille` (projets supprimés)
- Table `corbeille_image` (images des projets supprimés)
- Table `corbeille_technologies` (technologies supprimées)

### 3. Installation des dépendances Go

```bash
go mod download
```

## Lancement

Depuis le répertoire du projet :

```bash
go run .
```

Le serveur démarre sur `http://localhost:8080`

## Endpoints principaux

- `GET /` - Page d'accueil du portfolio
- `GET /admin` - Interface d'administration
- `GET /projet/{id}` - Détails d'un projet
- `GET /corbeille` - Gestion de la corbeille
- `POST /add-project` - Ajouter un projet
- `PUT /update-project` - Mettre à jour un projet
- `DELETE /delete-project/{id}` - Supprimer un projet

## Structure de données

### Project
```go
type Project struct {
    ID           int      `json:"id"`
    Titre        string   `json:"title"`
    DateCreation string   `json:"date"`
    Description  string   `json:"description"`
    Technologie  string   `json:"technologie"`
    Explication  string   `json:"explication"`
    Probleme     string   `json:"probleme"`
    Solution     string   `json:"solution"`
    UrlSource    string   `json:"url_source"`
    Images       []string `json:"images"`
}
```

## Développement

### Conventions

- Le code utilise le français pour les noms de variables et commentaires
- Les handlers HTTP sont organisés par domaine fonctionnel
- La couche MariaDB est isolée dans un package dédié
- Les structures de données sont centralisées dans `utils/struct.go`

### Configuration

La connexion à la base de données doit être configurée dans `mariadb/connection.go` avec les identifiants appropriés.

## Auteur

MATTEO-mart

## Licence

Ce projet est à usage personnel.
