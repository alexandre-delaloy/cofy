package services

import (
	"strconv"

	"github.com/blyndusk/cofy/app/core"
	"github.com/blyndusk/cofy/app/helpers"
	"github.com/blyndusk/cofy/app/middlewares"
	"github.com/bwmarrin/discordgo"
)

func UserNotFoundHandler(s *discordgo.Session, m *discordgo.MessageCreate, user core.User) {
	if strconv.Itoa(user.DiscordId) != m.Author.ID {
		helpers.EmbedCreatingProfile(s, m)
		middlewares.CreateUser(m)
		newUser := middlewares.GetUser(s, m.Author.ID)
		discordUser, _ := s.User(strconv.Itoa(newUser.DiscordId))
		helpers.EmbedViewProfile(s, m, newUser, discordUser)
	}
}
