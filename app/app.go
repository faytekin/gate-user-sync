package app

import (
	"fmt"
	"github.com/joho/godotenv"
	"macellan-gate-user-sync/alternatif"
	"macellan-gate-user-sync/helper"
	"macellan-gate-user-sync/ik"
	"os"
)

func Init() {
	_ = godotenv.Load(".env")
}

func SyncPersons() error {
	ikActivePhones, err := kolayIkPersons("1")
	if err != nil {
		helper.Log("Kolay IK active phone list failed", err)

		return err
	}

	ikPassivePhones, err := kolayIkPersons("0")
	if err != nil {
		helper.Log("Kolay IK passive phone list failed", err)

		return err
	}

	alternatifPhones, err := alternatifUsers()
	if err != nil {
		helper.Log("Alternatif user phone list failed", err)

		return err
	}

	willRemove := helper.FindToBeRemoved(ikPassivePhones, alternatifPhones)
	willAdd := helper.FindToBeAdded(ikActivePhones, alternatifPhones)

	if len(willAdd) > 0 {
		err = alternatif.AddNewUsers(willAdd)
		if err != nil {
			helper.Log("Error while adding new users", err)
		}
	}

	if len(willRemove) > 0 {
		err = alternatif.RemoveUsers(willRemove)
		if err != nil {
			helper.Log("Error while removing users", err)
		}
	}

	return nil
}

func kolayIkPersons(status string) ([]string, error) {
	var logName string
	if status == "1" {
		logName = "ACTIVE"
	} else {
		logName = "PASSIVE"
	}

	helper.Log("Getting Kolay IK " + logName + " phone list...")

	phones, err := ik.GetPhoneList(status)
	if err != nil {
		return nil, fmt.Errorf("Kolay IK "+logName+" phone List Failed %w", err)
	}

	helper.Log("Total Kolay IK "+logName+" phone count => ", len(phones))

	return phones, nil
}

func alternatifUsers() ([]string, error) {
	users, err := alternatif.GetUserList(os.Getenv("ALTERNATIF_USER_GROUP_ID"))
	if err != nil {
		return nil, fmt.Errorf("alternatif user list failed %w", err)
	}

	var phoneNumbers []string
	for _, user := range users {
		phoneNumbers = append(phoneNumbers, user.Phone)
	}

	helper.Log("Total Alternatif user count => ", len(phoneNumbers))

	return phoneNumbers, nil
}
