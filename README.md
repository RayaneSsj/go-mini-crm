# Mini CRM

Un petit CRM en ligne de commande pour gérer des contacts.

## Comment l'utiliser

Lance le programme :
```bash
go run main.go
```

Tu vas voir un menu avec plusieurs options :
- Ajouter un contact
- Lister tous les contacts
- Supprimer un contact
- Mettre à jour un contact
- Quitter

## Ajouter un contact avec des flags

Tu peux aussi ajouter un contact directement en ligne de commande :
```bash
go run main.go -nom "Jean Dupont" -email "jean@email.com"
```

Ou avec un ID spécifique :
```bash
go run main.go -id 5 -nom "Marie Martin" -email "marie@email.com"
```

## Build

Pour compiler le programme :
```bash
go build
```

Ensuite tu peux l'exécuter avec `./go-mini-crm` (ou `go-mini-crm.exe` sur Windows).