package app

import (
	"fmt"
	"gate-user-sync/helper"
	"gate-user-sync/ik"
	"gate-user-sync/macellan"
)

func Init() error {
	var err error
	var config *helper.Config
	config, err = helper.GetConfig()
	if err != nil {
		return fmt.Errorf("failed to get config: %w", err)
	}

	err = helper.CheckConfig(config)
	if err != nil {
		return err
	}

	ik.SetConfig(config)
	macellan.SetConfig(config)

	return nil
}

func SyncPersons() error {
	ikActivePhones, err := getIkActivePhones()
	if err != nil {
		return err
	}

	ikPassivePhones, err := getIkPassivePhones()
	if err != nil {
		return err
	}

	macellanPhones, err := getMacellanPhones()
	if err != nil {
		return err
	}

	addPhones := helper.FindToBeAdded(ikActivePhones, macellanPhones)
	removePhones := helper.FindToBeRemoved(ikPassivePhones, macellanPhones)
	addPhonesLength := len(addPhones)
	removePhonesLength := len(removePhones)

	if addPhonesLength > 0 {
		err = macellan.AddNewUsers(addPhones)
		if err != nil {
			helper.Log("Macellan error while adding new users", err)
		}
	}

	if removePhonesLength > 0 {
		err = macellan.RemoveUsers(removePhones)
		if err != nil {
			helper.Log("Macellan error while removing users", err)
		}
	}

	if addPhonesLength == 0 && removePhonesLength == 0 {
		helper.Log("No changes. The lists are already up to date.")
	}

	return nil
}

func getIkActivePhones() ([]string, error) {
	helper.Log("Getting IK active phone list...")

	phones, err := ik.GetPhoneList("1")
	if err != nil {
		return nil, err
	}

	helper.Log("Total IK active phone count => ", len(phones))

	return phones, nil
}

func getIkPassivePhones() ([]string, error) {
	helper.Log("Getting IK passive phone list...")

	phones, err := ik.GetPhoneList("0")
	if err != nil {
		return nil, err
	}

	helper.Log("Total IK passive phone count => ", len(phones))

	return phones, nil
}

func getMacellanPhones() (map[string]bool, error) {
	phones, err := macellan.GetPhoneList()
	if err != nil {
		return nil, err
	}

	helper.Log("Total Macellan user count => ", len(phones))

	return phones, nil
}
