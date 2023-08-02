package app

import (
	"fmt"
	"gate-user-sync/alternatif"
	"gate-user-sync/helper"
	"gate-user-sync/ik"
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
	alternatif.SetConfig(config)

	return nil
}

func SyncPersons() error {
	ikActivePhones, err := getIkActivePersons()
	if err != nil {
		return err
	}

	ikPassivePhones, err := getIkPassivePersons()
	if err != nil {
		return err
	}

	alternatifPhones, err := alternatifUsers()
	if err != nil {
		return err
	}

	addPhones := helper.FindToBeAdded(ikActivePhones, alternatifPhones)
	removePhones := helper.FindToBeRemoved(ikPassivePhones, alternatifPhones)

	if len(addPhones) > 0 {
		err = alternatif.AddNewUsers(addPhones)
		if err != nil {
			helper.Log("Alternatif error while adding new users", err)
		}
	}

	if len(removePhones) > 0 {
		err = alternatif.RemoveUsers(removePhones)
		if err != nil {
			helper.Log("Alternatif error while removing users", err)
		}
	}

	return nil
}

func getIkActivePersons() ([]string, error) {
	helper.Log("Getting IK active phone list...")

	phones, err := ik.GetPhoneList("1")
	if err != nil {
		return nil, err
	}

	helper.Log("Total IK active phone count => ", len(phones))

	return phones, nil
}

func getIkPassivePersons() ([]string, error) {
	helper.Log("Getting IK passive phone list...")

	phones, err := ik.GetPhoneList("0")
	if err != nil {
		return nil, err
	}

	helper.Log("Total IK passive phone count => ", len(phones))

	return phones, nil
}

func alternatifUsers() ([]string, error) {
	phones, err := alternatif.GetPhoneList()
	if err != nil {
		return nil, err
	}

	helper.Log("Total Alternatif user count => ", len(phones))

	return phones, nil
}
