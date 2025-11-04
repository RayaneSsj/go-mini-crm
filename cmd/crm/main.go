package main

import (
	"go-mini-crm/internal/app"
	"go-mini-crm/internal/storage"
)

func main() {
	store := storage.NewMemoryStore()
	app.Run(store)
}
