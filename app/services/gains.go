package services

import (
	"math"

	"github.com/blyndusk/cofy/app/core"
	"github.com/blyndusk/cofy/app/helpers"
	"github.com/blyndusk/cofy/app/middlewares"
	"github.com/bwmarrin/discordgo"
)

func HandleGains(s *discordgo.Session, m *discordgo.MessageCreate, user core.User) {

	gains := core.Gains{
		Coins: int(math.Round(float64(len(m.Content)) / 10)),
		Xp:    int(math.Round(float64(len(m.Content)) / 5)),
	}
	// update user values
	middlewares.UpdateGains(s, m, user, gains)
	// display gains
	helpers.EmbedViewGains(s, m, gains)
}
