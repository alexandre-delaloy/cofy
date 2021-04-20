package services

import (
	"os"

	"github.com/blyndusk/cofy/app/helpers"
	log "github.com/sirupsen/logrus"

	"github.com/joho/godotenv"
)

func EnvVar(key string) string {
	err := godotenv.Load(".env")
	helpers.ExitOnError("Error loading .env file", err)

	envVars, doesExist := os.LookupEnv(key)
	if !doesExist {
		log.Error("Couldn't find variable : ", key)
	}

	return envVars
}
