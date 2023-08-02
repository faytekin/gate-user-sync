package main

import (
	"gate-user-sync/app"
	"log"
)

func main() {
	err := app.Init()
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}

	err = app.SyncPersons()
	if err != nil {
		log.Fatalf("Failed to sync persons: %v", err)
	}
}
