package middlewares

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/blyndusk/cofy/app/core"
	"github.com/blyndusk/cofy/app/helpers"
	"github.com/bwmarrin/discordgo"
)

func GetDrinks(s *discordgo.Session) core.Drinks {
	// setup url
	url := fmt.Sprintf("%s/drinks", helpers.EnvVar("API_URL"))
	// declare user
	var drinks core.Drinks
	// get the user
	resp, err := http.Get(url)
	helpers.ExitOnError("Error while getting user", err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	helpers.ExitOnError("Error while reading request body", err)
	json.Unmarshal(body, &drinks)
	return drinks
}