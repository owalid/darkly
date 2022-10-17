on peux voir dans les cookies: 68934a3e9455fa72420237eb05902327 qui quand on le dechiffre en md5 on obtiens la valeur "false".

Il suffit de chiffré en md5 la chaine "true" pour passer en admin.

true => "b326b5062b2f0e69046810717534cb09"


### Amelioration

Utiliser un algo de hashage plus sécurisé que md5, qui utilise un salt pour chiffré la chaine.
Ou utiliser des tokens type jwt avec une clef de chiffrement AES256.