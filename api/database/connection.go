package database

import (
	"fmt"
	"time"

	"github.com/blyndusk/cofy/api/services"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Db is the database object
var Db *gorm.DB

// Config is the structure used to load db credentials from the environment.
type Config struct {
	Name     string `env:"DB_NAME"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	Host     string `env:"DB_HOST"`
	Port     int    `env:"DB_PORT" envDefault:"5432"`
}

// Init Initializes a db connection
func Init(cfg Config) error {

	dbURL := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		services.SetupEnv("DB_HOST"),
		services.SetupEnv("DB_USER"),
		services.SetupEnv("DB_PASSWORD"),
		services.SetupEnv("DB_NAME"),
		services.SetupEnv("DB_PORT"),
	)
	log.Info(dbURL)
	var tmpDb *gorm.DB
	var err error

	// Try connecting database 5 times
	for test := 1; test <= 5; test++ {
		tmpDb, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})

		if err != nil {
			log.Warn("db connection failed. (%s/5)", test)
			time.Sleep(5 * time.Second)
			continue
		}

		break
	}
	if err != nil {
		return err
	}

	Db = tmpDb
	log.Info("Connected to database!")

	return nil
}
