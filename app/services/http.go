package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type User struct {
	Id        uint   `json:"id"`
	DiscordId string `json:"discord_id"`
	Name      string `json:"name"`
	Coins     int    `json:"coins"`
	Xp        int    `json:"xp"`
}

func GetUser(discordId string) User {
	var user User

	resp, err := http.Get(fmt.Sprintf("%s/users/d/%s", EnvVar("API_URL"), discordId))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	json.Unmarshal(body, &user)
	log.Info(user.Coins)
	if err != nil {
		log.Fatal(err)
	}
	return user
}
