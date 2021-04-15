package commands

import (
	"github.com/bwmarrin/discordgo"
)

func info(s *discordgo.Session, m *discordgo.MessageCreate) {
	infoCmd := Command{
		"info",
		[]string{"info", "i"},
		"Give informatiosn about Cofy bot.",
		"Hi! I'm Cofy, a Discord bot made with the Golang programming language.\nMy purpose is to help you to gain cOoOOooOoooffEeEEeeEeee...",
	}

	if matchCommand(false, s, m, infoCmd.aliases) == true {
		callExecute(infoCmd, s, m)
	}
	if matchCommand(true, s, m, infoCmd.aliases) == true {
		callHelp(infoCmd, s, m)
	}
}

func (command Command) Execute(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, command.stringResponse)
}

func (command Command) Help(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, command.description)
}
