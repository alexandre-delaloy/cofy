package services

import (
	"math"

	"github.com/blyndusk/cofy/app/core"
	"github.com/blyndusk/cofy/app/helpers"
	"github.com/blyndusk/cofy/app/middlewares"
	"github.com/bwmarrin/discordgo"
)

func GainsHandler(s *discordgo.Session, m *discordgo.MessageCreate, user core.User) {
	gainedCoins := int(math.Round(float64(len(m.Content)) / 10))
	gainedXp := int(math.Round(float64(len(m.Content)) / 5))
	middlewares.UpdateGains(s, m, user, gainedCoins, gainedXp)
	helpers.EmbedViewGains(s, m, gainedCoins, gainedXp)
}
