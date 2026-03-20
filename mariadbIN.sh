#!/bin/bash
clear 

if [ "$EUID" -ne 0 ]; then
    echo "Le fichier est à lancer en tant que root"
    exit 1
fi
echo "----------------------------------------------"
echo "          MAJ"
echo "----------------------------------------------"
zypper refresh
zypper dup -y

echo "----------------------------------------------"
echo "          INSTALLATION DE MARIADB"
echo "----------------------------------------------"
zypper install -y mariadb

echo "----------------------------------------------"
echo "          MAJ"
echo "----------------------------------------------"
zypper refresh
zypper dup -y

echo "----------------------------------------------"
echo "          Lancement de MariaDB"
echo "----------------------------------------------"
systemctl enable mariadb
systemctl start mariadb

echo "-----------------------------------------------------"
echo "          Connexion à MariaDB via le compte matteo"
echo "-----------------------------------------------------"
mariadb -u matteo -p
