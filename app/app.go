package app

import (
	"github.com/joho/godotenv"
)

func Init() {
	_ = godotenv.Load(".env")
}

func SyncPersons() {
	activePersonPhones, err := getKolayIkPersonPhoneList("1")

	if err != nil {
		Log("Kolay IK PersonIds List Failed", err)
	}

	Log("Total Kolay IK active phone count => ", len(activePersonPhones))
}
