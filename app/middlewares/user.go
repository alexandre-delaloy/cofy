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
	// setup url
	url := fmt.Sprintf("%s/users/%s", core.ApiUrl, userId)
	// declare user
	var user core.User
	// get the user
	resp, err := http.Get(url)
	helpers.ExitOnError("Error while getting user", err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	helpers.ExitOnError("Error while reading request body", err)
	json.Unmarshal(body, &user)
	return user
}

func CreateUser(discordUser *discordgo.User) {
	// setup url
	url := fmt.Sprintf("%s/users", core.ApiUrl)
	// get the string discord id
	discordId, _ := strconv.Atoi(discordUser.ID)
	// setup the payload
	payload := map[string]interface{}{
		"discord_id": discordId,
		"name":       discordUser.Username,
		"coins":      0,
		"xp":         0,
	}
	data, err := json.Marshal(payload)
	helpers.ExitOnError("Error while parsing json", err)
	// post the user
	resp, err := http.Post(url, "application/json",
		bytes.NewBuffer(data))
	helpers.ExitOnError("Error while posting user", err)
	// get user response
	var res map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&res)
}
