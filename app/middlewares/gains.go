package middlewares

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/blyndusk/cofy/app/core"
	"github.com/blyndusk/cofy/app/helpers"
	"github.com/bwmarrin/discordgo"
)

func UpdateGains(s *discordgo.Session, m *discordgo.MessageCreate, user core.User, gainedCoins int, gainedXp int) {
	// setup url
	url := fmt.Sprintf("%s/users/%s", helpers.EnvVar("API_URL"), m.Author.ID)
	// setup payload
	payload, err := json.Marshal(map[string]interface{}{
		"coins": user.Coins + gainedCoins,
		"xp":    user.Xp + gainedXp,
	})
	helpers.ExitOnError("Error while parsing json", err)
	// setup client
	client := &http.Client{}
	// put gains
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(payload))
	helpers.ExitOnError("Error while putting user", err)
	// setup header
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	helpers.ExitOnError("Error while doing request", err)

	defer resp.Body.Close()

}
