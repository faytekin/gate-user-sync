package app

import (
	"github.com/joho/godotenv"
)

func Init() {
	_ = godotenv.Load(".env")
}

func SyncPersons() {
	Log("Getting Kolay IK ACTIVE phone list...")

	activePhones, err := getKolayIkPersonPhoneList("1")
	if err != nil {
		Log("Kolay IK active phone List Failed", err)
	}

	Log("Total Kolay IK active phone count => ", len(activePhones))

	Log("Getting Kolay IK PASSIVE phone list...")

	passivePhones, err := getKolayIkPersonPhoneList("0")
	if err != nil {
		Log("Kolay IK passive phone list failed", err)
	}

	Log("Total Kolay IK passive phone count => ", len(passivePhones))
}
