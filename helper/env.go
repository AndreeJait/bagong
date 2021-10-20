package helper

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetBaseBagong() (string, error) {
	baseFile := os.Getenv("BASE_BAGONG")

	if baseFile == "" {
		log.Println("Base bagong not found")
		return "", errors.New("base bagong not found")
	}
	return baseFile, nil
}

func goDotEnvVariable(key string) (string, error) {

	// load .env file
	baseFile := os.Getenv("BASE_BAGONG")
	err := godotenv.Load(baseFile + "/.env")

	if err != nil {
		return "", fmt.Errorf("sorry can't load env \nerror: %s", err.Error())
	}

	return os.Getenv(key), nil
}
