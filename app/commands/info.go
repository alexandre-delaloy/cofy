package commands

import (
	"github.com/bwmarrin/discordgo"
)

func Info(s *discordgo.Session, m *discordgo.MessageCreate) {

	infoCmd := Command{
		"info",
		[]string{"info", "i"},
		"Give informatiosn about Cofy bot.",
		"Hi! I'm Cofy, a Discord bot made with the Golang programming language.\nMy purpose is to help you to gain cOoOOooOoooffEeEEeeEeee...",
	}
	UseCommand(s, m, infoCmd)

}
