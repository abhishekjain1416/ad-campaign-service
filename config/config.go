package config

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	res := strings.Split(pwd, "ad-campaign-service")
	errenv := godotenv.Load(res[0] + "ad-campaign-service/.env")
	if errenv != nil {
		log.Println(pwd)
		log.Fatal(errenv)
	}
}
