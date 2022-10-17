Sur les liens en bas de la page index on observe un params: page=redirect&site

la page redirect prend un query params `site` qui permet de rediriger vers une url.

Nous pouvons alors changer le site vers un site externe.

payload:
```
http://172.16.42.21/index.php?page=redirect&site=https://google.com
```

### Amelioration

Faire en sorte d'etablir une white list ou authoriser uniquement les pages du domaine courrant.