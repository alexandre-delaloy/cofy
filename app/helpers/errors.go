package helpers

import (
	log "github.com/sirupsen/logrus"
)

func ExitOnError(msg string, err error) {
	if err != nil {
		log.Fatal(msg, err)
	}
}
