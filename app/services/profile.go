package services

import (
	"github.com/blyndusk/cofy/app/core"
	"github.com/bwmarrin/discordgo"
)

func CreateUserIfDoesntExist(user core.User, s *discordgo.Session, m *discordgo.MessageCreate) {

	if user.DiscordId != m.Author.ID {
		EmbedProfileNotFound(s, m)
		EmbedProfileCreating(s, m)
		CreateUser(m)
		EmbedProfileCreated(s, m)
	}
}
