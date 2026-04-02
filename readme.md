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
