package services

import (
	"strconv"

	"github.com/blyndusk/cofy/app/core"
	"github.com/bwmarrin/discordgo"
)

func CreateUserIfDoesntExist(user core.User, s *discordgo.Session, m *discordgo.MessageCreate) {

	if strconv.Itoa(int(user.DiscordId)) != m.Author.ID {
		EmbedProfileNotFound(s, m)
		EmbedProfileCreating(s, m)
		CreateUser(m)
		EmbedProfileCreated(s, m)
	}
}
