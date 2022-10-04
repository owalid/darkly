Sur les liens en bas de la page index on observe un params: page=redirect&site

si on change le site en n'importe quoi, on obtiens le flag

payload:
```
http://172.16.42.21/index.php?page=redirect&site=https://google.com
```