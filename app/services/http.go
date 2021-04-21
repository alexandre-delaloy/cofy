package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

type User struct {
	Id        uint      `json:"id"`
	DiscordId string    `json:"discord_id"`
	Name      string    `json:"name"`
	Coins     int       `json:"coins"`
	Xp        int       `json:"xp"`
	CreatedAt time.Time `json:"created_at"`
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
	log.Info(user)
	if err != nil {
		log.Fatal(err)
	}
	return user
}

func PostUser(discordUser *discordgo.User) {
	type PostUser struct {
		DiscordId string `json:"discord_id"`
		Name      string `json:"name"`
		Coins     int    `json:"coins"`
		Xp        int    `json:"xp"`
	}    
	values := map[string]interface{}{
		"discord_id": discordUser.ID,
		"name":       discordUser.Username,
		"coins":      0,
		"xp":         0,
	}
	json_data, err := json.Marshal(values)
	resp, err := http.Post(fmt.Sprintf("%s/users", EnvVar("API_URL")),  "application/json",
	bytes.NewBuffer(json_data))
	if err != nil {
		log.Fatal(err)
	}

	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)

	fmt.Println(res["json"])
}
