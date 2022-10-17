On peu voir que certaines pages sont charger avec un query parameters `page`:

> http://192.168.99.102/?page=<page_name>


Nous pouvons alors essayer de lancer une attaque `transversal path`


payload:

```
../../../../../../../../../etc/passwd
```

On obtiens le flag :)

Amelioration:

Utiliser un chemin absolu pour les images, et ne pas utiliser de chemin relatif.
Proteger le parametre page pour pouvoir utiliser uniquement des fichiers php, et ne pas pouvoir remonter dans le dossier.