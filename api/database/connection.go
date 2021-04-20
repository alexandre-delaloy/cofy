package database

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/blyndusk/cofy/api/helpers"
	"github.com/blyndusk/cofy/api/services"
)

var Db *gorm.DB

func Connect() error {
	dbURL := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		services.EnvVar("DB_HOST"),
		services.EnvVar("DB_USER"),
		services.EnvVar("DB_PASSWORD"),
		services.EnvVar("DB_NAME"),
		services.EnvVar("DB_PORT"),
	)
	var tmpDb *gorm.DB
	var err error

	for i := 1; i <= 5; i++ {
		tmpDb, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})
		if err != nil {
			log.Warn(fmt.Sprintf("Failed to connect to database. Retry... (%d/5)", i))
			time.Sleep(5 * time.Second)
			continue
		}
		break
	}
	helpers.ExitOnError("Failed to connecto to database", err)

	Db = tmpDb
	log.Info("Connected to database !")
	return nil
}
