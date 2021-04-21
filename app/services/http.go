package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/blyndusk/cofy/app/core"
	"github.com/blyndusk/cofy/app/helpers"
	"github.com/bwmarrin/discordgo"
)

func GetUser(discordId string) core.User {
	var user core.User

	resp, err := http.Get(fmt.Sprintf("%s/users/%s", EnvVar("API_URL"), discordId))
	helpers.ExitOnError("Error while getting user", err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	helpers.ExitOnError("Error while reading request body", err)

	json.Unmarshal(body, &user)
	return user
}

func CreateUser(m *discordgo.MessageCreate) {
	userId, err := strconv.Atoi(m.Author.ID)
	values := map[string]interface{}{
		"discord_id": userId,
		"name":       m.Author.Username,
		"coins":      0,
		"xp":         0,
	}
	json_data, err := json.Marshal(values)
	helpers.ExitOnError("Error while parsing json", err)

	resp, err := http.Post(fmt.Sprintf("%s/users", EnvVar("API_URL")), "application/json",
		bytes.NewBuffer(json_data))
	helpers.ExitOnError("Error while posting user", err)

	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)
}

func UpdateGains(s *discordgo.Session, m *discordgo.MessageCreate, gainedCoins int, gainedXp int) {
	user := GetUser(m.Author.ID)
	CreateUserIfDoesntExist(user, s, m)

	payload, err := json.Marshal(map[string]interface{}{
		"coins": user.Coins + gainedCoins,
		"xp":    user.Xp + gainedXp,
	})
	helpers.ExitOnError("Error while parsing json", err)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/users/%s", EnvVar("API_URL"), m.Author.ID), bytes.NewBuffer(payload))
	helpers.ExitOnError("Error while putting user", err)

	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	helpers.ExitOnError("Error while doing request", err)

	defer resp.Body.Close()
}
