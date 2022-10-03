On peux voir qu'une image est chargé de cette facon: background-image: "../images/banner.jpg"
ce qui peux nous alerter sur l'exploit transversal path

On peu tester une attaque LFI: ../../../../../../../../../etc/passwd

On obtiens le flag :)

Amelioration:

Utiliser un chemin absolu pour les images, et ne pas utiliser de chemin relatif.
Proteger le parametre page pour pouvoir utiliser uniquement des fichiers php, et ne pas pouvoir remonter dans le dossier.