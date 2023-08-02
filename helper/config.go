package helper

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	IKApi         IKAPI
	AlternatifApi AlternatifAPI
}

type IKAPI struct {
	Url         string
	BearerToken string
}

type AlternatifAPI struct {
	Url         string
	BearerToken string
	GroupId     string
}

func GetConfig() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return &Config{}, err
	}

	return &Config{
		IKApi: IKAPI{
			Url:         os.Getenv("KOLAY_IK_API_URL"),
			BearerToken: os.Getenv("KOLAY_IK_API_TOKEN"),
		},
		AlternatifApi: AlternatifAPI{
			Url:         os.Getenv("ALTERNATIF_API_URL"),
			BearerToken: os.Getenv("ALTERNATIF_API_TOKEN"),
			GroupId:     os.Getenv("ALTERNATIF_USER_GROUP_ID"),
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

	if config.AlternatifApi.Url == "" {
		return fmt.Errorf("please put ALTERNATIF_API_URL to .env file")
	}

	if config.AlternatifApi.BearerToken == "" {
		return fmt.Errorf("please put ALTERNATIF_API_TOKEN to .env file")
	}

	if config.AlternatifApi.BearerToken == "" {
		return fmt.Errorf("please put ALTERNATIF_API_TOKEN to .env file")
	}

	return nil
}
