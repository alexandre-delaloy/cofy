package middlewares

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"net/http"

	"github.com/blyndusk/cofy/app/core"
	"github.com/blyndusk/cofy/app/helpers"
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

func UpdateGains(s *discordgo.Session, m *discordgo.MessageCreate, user core.User, gains core.Gains) {
	newGains := core.Gains{
		Coins: user.Coins + gains.Coins,
		Xp:    user.Xp + gains.Xp,
		Level: int(0.1666 * math.Sqrt(float64(user.Xp+gains.Xp))),
	}
	logrus.Info("new gains")
	logrus.Info(newGains)
	// setup url
	url := fmt.Sprintf("%s/users/%s", helpers.EnvVar("API_URL"), m.Author.ID)
	// setup payload
	payload, err := json.Marshal(map[string]interface{}{
		"coins": newGains.Coins,
		"xp":    newGains.Xp,
		"level": newGains.Level,
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
	logrus.Info(bytes.NewBuffer(payload))
	helpers.ExitOnError("Error while doing request", err)

	defer resp.Body.Close()

}
