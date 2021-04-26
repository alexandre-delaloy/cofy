package services

import (
	"math"

	"github.com/blyndusk/cofy/app/core"
	"github.com/blyndusk/cofy/app/helpers"
	"github.com/blyndusk/cofy/app/middlewares"
	"github.com/bwmarrin/discordgo"
)

func HandleGains(s *discordgo.Session, m *discordgo.MessageCreate, user core.User) {
	// calculate gained coins according to message length
	gainedCoins := int(math.Round(float64(len(m.Content)) / 10))
	// calculate gained xp  according to message length
	gainedXp := int(math.Round(float64(len(m.Content)) / 5))
	// update user values
	middlewares.UpdateGains(s, m, user, gainedCoins, gainedXp)
	// display gains
	helpers.EmbedViewGains(s, m, gainedCoins, gainedXp)
}
