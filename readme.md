## Présentation du projet

Ce projet est bien plus qu'un simple CV en ligne. Il s'agit d'une application web dynamique développée en Go (Golang), conçue pour gérer et afficher mes projets et compétences de manière automatisée.

L'application communique avec une base de données MariaDB pour le stockage des données et utilise le moteur de templates natif de Go pour le rendu côté serveur.
## Architecture Technique

Le projet suit une structure logique conforme aux bonnes pratiques enseignées en BTS SIO :

    Backend : Go (Langage performant et robuste).

    Base de données : MariaDB (Gestion du patrimoine de données).

    Frontend : HTML5, CSS3, JavaScript (Interface utilisateur).

    Automatisation : Scripts Shell (.sh) pour l'initialisation de l'environnement.

## Compétences SIO (SLAM) validées

    B1.1 - Gérer le patrimoine informatique : Utilisation de Git et structuration d'un projet complexe.

    B1.3 - Développer la présence en ligne : Création d'une application web complète.

    Problématique SIO : Comment automatiser le déploiement d'une base de données locale pour une application de production ?

## Structure du dépôt
    Dossier / Fichier	Rôle
    main.go	Point d'entrée, configuration des routes et du serveur.
    /templates	Fichiers HTML dynamiques.
    /route	Gestion des différentes pages et de la logique de navigation.
    /mariadb	Scripts SQL et logique liée à la persistance des données.
    /utils	Fonctions utilitaires réutilisables.
    mariadbIN.sh	Script d'automatisation pour configurer MariaDB rapidement.
## Installation et Lancement

Pour déployer le projet sur un environnement de développement local :
### 1. Préparer la base de données

Le script mariadbIN.sh automatise la création des tables et des droits.
    Bash

    chmod +x mariadbIN.sh
    sudo ./mariadbIN.sh

### 2. Lancer l'application

Assurez-vous d'avoir Go installé, puis lancez le serveur :
    Bash
    
    go run .
