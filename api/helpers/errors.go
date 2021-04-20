package helpers

import (
	log "github.com/sirupsen/logrus"
)

func Error(msg string, err error) {
	if err != nil {
		log.Error(msg, err)
	}
}

func ExitOnError(msg string, err error) {
	if err != nil {
		log.Fatal(msg, err)
	}
}
