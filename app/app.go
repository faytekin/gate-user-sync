package app

import (
	"fmt"
	"github.com/joho/godotenv"
)

func Init() {
	_ = godotenv.Load(".env")
}

func SyncPersons() {
	activePersons, err := GetKolayIKPersonList("1")

	if err != nil {
		Log("Kolay IK PersonIds List Failed", err)
	}

	Log("Total Kolay IK active personal count => ", len(activePersons))

	var phoneNumbers []string
	for _, person := range activePersons {
		formattedPhoneNumber := person.GetFormattedPhoneNumber()

		if formattedPhoneNumber != "" {
			phoneNumbers = append(phoneNumbers, formattedPhoneNumber)
		}
	}

	for _, person := range activePersons {
		phoneNumber := person.GetFormattedPhoneNumber()

		if phoneNumber == "" || len(phoneNumber) < 11 {
			fmt.Println("Name:", person.FirstName, person.LastName)
			fmt.Println("Phone:", phoneNumber)
			fmt.Println("--------------------------------")
		}
	}
}
