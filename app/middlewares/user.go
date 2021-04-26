package middlewares

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

func GetUser(s *discordgo.Session, userId string) core.User {
	url := fmt.Sprintf("%s/users/%s", helpers.EnvVar("API_URL"), userId)

	var user core.User
	resp, err := http.Get(url)
	helpers.ExitOnError("Error while getting user", err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	helpers.ExitOnError("Error while reading request body", err)

	json.Unmarshal(body, &user)

	return user
}

func CreateUser(m *discordgo.MessageCreate) {
	url := fmt.Sprintf("%s/users", helpers.EnvVar("API_URL"))

	userId, err := strconv.Atoi(m.Author.ID)
	values := map[string]interface{}{
		"discord_id": userId,
		"name":       m.Author.Username,
		"coins":      0,
		"xp":         0,
	}
	data, err := json.Marshal(values)
	helpers.ExitOnError("Error while parsing json", err)

	resp, err := http.Post(url, "application/json",
		bytes.NewBuffer(data))
	helpers.ExitOnError("Error while posting user", err)

	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)
}
