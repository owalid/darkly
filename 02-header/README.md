Dans ?page=b7e44c7a40c5f80139f0a50f3650fb2bd8d00b0d24667c4c2ca32c88e13b758f (lien BornToSec)

on peux voir un commentaire super long quand on ouvre le code source de la page

on en deduis alors qu'il faut venir du site de la nsa et de changer l'user agent en "ft_bornToSec"


> Let's use this browser : "ft_bornToSec"
> You must come from : "https://www.nsa.gov/"

Changer les headers pour:

```
Referer: "https://www.nsa.gov/"
User-Agent: "ft_bornToSec"
```

