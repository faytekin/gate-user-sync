package main

import (
	"gate-user-sync/app"
	"time"
)

func main() {
	app.Init()

	_ = app.SyncPersons()

	ticker := time.NewTicker(1 * time.Minute)

	for range ticker.C {
		_ = app.SyncPersons()
	}
}
