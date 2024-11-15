package helper

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	IKApi       IKAPI
	MacellanApi MacellanAPI
}

type IKAPI struct {
	Url         string
	BearerToken string
}

type MacellanAPI struct {
	Url         string
	BearerToken string
	GroupId     string
}

func GetConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return &Config{}, err
	}

	return &Config{
		IKApi: IKAPI{
			Url:         os.Getenv("KOLAY_IK_API_URL"),
			BearerToken: os.Getenv("KOLAY_IK_API_TOKEN"),
		},
		MacellanApi: MacellanAPI{
			Url:         os.Getenv("MACELLAN_API_URL"),
			BearerToken: os.Getenv("MACELLAN_API_TOKEN"),
			GroupId:     os.Getenv("MACELLAN_USER_GROUP_ID"),
		},
	}, nil
}

func CheckConfig(config *Config) error {
	if config.IKApi.Url == "" {
		return fmt.Errorf("please put KOLAY_IK_API_URL to .env file")
	}

	if config.IKApi.BearerToken == "" {
		return fmt.Errorf("please put KOLAY_IK_API_TOKEN to .env file")
	}

	if config.MacellanApi.Url == "" {
		return fmt.Errorf("please put MACELLAN_API_URL to .env file")
	}

	if config.MacellanApi.BearerToken == "" {
		return fmt.Errorf("please put MACELLAN_API_TOKEN to .env file")
	}

	if config.MacellanApi.BearerToken == "" {
		return fmt.Errorf("please put MACELLAN_API_TOKEN to .env file")
	}

	return nil
}
