package commands

import (
	"github.com/bwmarrin/discordgo"
)

var Prefix string = "c!"

type ICommand interface {
	Execute(s *discordgo.Session, m *discordgo.MessageCreate)
	Help(s *discordgo.Session, m *discordgo.MessageCreate)
}

type Command struct {
	name           string
	aliases        []string
	description    string
	stringResponse string
}

func SetCommands(s *discordgo.Session, m *discordgo.MessageCreate) {
	Info(s, m)
}
