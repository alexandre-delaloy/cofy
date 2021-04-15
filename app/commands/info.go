package commands

import (
	"github.com/bwmarrin/discordgo"
)


func (command Command) Execute(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, command.stringResponse)
}

func callExecute(i ICommand, s *discordgo.Session, m *discordgo.MessageCreate) {
	i.Execute(s, m)
}

func (command Command) Help(s *discordgo.Session, m *discordgo.MessageCreate) {
	// thing to execute
	s.ChannelMessageSend(m.ChannelID, command.description)

}

func callHelp(i ICommand, s *discordgo.Session, m *discordgo.MessageCreate) {
	i.Help(s, m)
}

func Info(s *discordgo.Session, m *discordgo.MessageCreate) {

	infoCmd := Command{
		"info",
		[]string{"info", "i"},
		"Give informatiosn about Cofy bot.",
		"Hi! I'm Cofy, a Discord bot made with the Golang programming language.\nMy purpose is to help you to gain cOoOOooOoooffEeEEeeEeee...",
	}

	// UseCommand(s, m, infoCmd)
	if MatchCommand(false, s, m, infoCmd.aliases) == true {
		callExecute(infoCmd, s, m)
	}
	if MatchCommand(true, s, m, infoCmd.aliases) == true {
		callHelp(infoCmd, s, m)
	}
}
