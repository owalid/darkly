Sur la page recover, l'email est inscrit en dur dans le html, il est possible de le modifier sur la page ou dans la requete directement.

avec un body: `mail=lol%40borntosec.com&Submit=Submit`

On peut donc envoyer un mail à l'adresse de notre choix. On peu egalement recuperer un mail interne "webmaster@borntosec.com".

### Amelioration

Gerer ca en back, si l'utilisateur n'existe pas en db, ne pas envoyer de mail