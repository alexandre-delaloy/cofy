package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func Info(s *discordgo.Session, m *discordgo.MessageCreate, prefix string) {
	command := "info"
	description := "Give informations about Cofy Bot"
	response := "Hi! I'm Cofy, a Discord bot made with the Golang programming language.\nMy purpose is to help you to gain cOoOOooOoooffEeEEeeEeee..."

	if m.Content == fmt.Sprintf("%s %s", prefix, command) {
		s.ChannelMessageSend(m.ChannelID, response)
	}

	if m.Content == fmt.Sprintf("%s help %s", prefix, command) {
		s.ChannelMessageSend(m.ChannelID, description)
	}
}
