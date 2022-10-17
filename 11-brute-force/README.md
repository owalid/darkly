Nous avons accees a une page de connexion. Aucune faille ne semble visible dans les requetes ou header http.

Allons chercher un petit dictionnaire sur: https://wiki.skullsecurity.org/Passwords#Password_dictionaries

L'idée est de faire une boucle sur chaque mot du dictionnaire et de tester si le mot de passe est le bon.


```
go run bruteforce.go
2022/10/03 14:05:42 Dictinary rockyou.txt does not exist
2022/10/03 14:05:42 Download dictonnary.....
2022/10/03 14:06:49 File downloaded
2022/10/03 14:06:49 Open dictonnary
2022/10/03 14:06:49 Start bruteforce....
2022/10/03 14:09:13 [+] Win : shadow  
```


### Amelioration:


Proteger le formulaire de login des attaques par force brute.

- Ajouter un captcha
- Ajouter un token csrf
- Limiter le nombre de tentatives
