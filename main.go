package main

import (
	"macellan-gate-user-sync/app"
	"time"
)

func main() {
	app.Init()

	app.SyncPersons()

	ticker := time.NewTicker(1 * time.Minute)

	for range ticker.C {
		app.SyncPersons()
	}
}
