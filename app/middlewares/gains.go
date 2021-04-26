package middlewares

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/blyndusk/cofy/app/core"
	"github.com/blyndusk/cofy/app/helpers"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

func UpdateGains(s *discordgo.Session, m *discordgo.MessageCreate, user core.User, gainedCoins int, gainedXp int) {
	url := fmt.Sprintf("%s/users/%s", helpers.EnvVar("API_URL"), m.Author.ID)

	payload, err := json.Marshal(map[string]interface{}{
		"coins": user.Coins + gainedCoins,
		"xp":    user.Xp + gainedXp,
	})
	log.Info(fmt.Sprintf("cf: +%d - xp: +%d", gainedCoins, gainedXp))
	helpers.ExitOnError("Error while parsing json", err)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(payload))
	helpers.ExitOnError("Error while putting user", err)

	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	helpers.ExitOnError("Error while doing request", err)

	defer resp.Body.Close()

}
