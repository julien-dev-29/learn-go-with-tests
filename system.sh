#! /bin/bash

while true; do
    clear

    echo "================================"
    echo "         SYSTEM TOOL"
    echo "================================"
    echo "1. Informations système"
    echo "2. Utilisation du disque"
    echo "3. Utilisation de la RAM"
    echo "4. Processus actifs"
    echo "5. Vérifier le réseau"
    echo "0. Quitter"
    echo "================================"

    read -p "Votre choix : " choix

    case $choix in
        1)
            uname -a
            ;;
        2)
            df -h
            ;;
    esac

    read -p "Appuyer sut entrer pour continuer..."
done