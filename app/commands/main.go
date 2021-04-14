package commands

import (
	"github.com/bwmarrin/discordgo"
)

var Prefix string = "c!"

type Command struct {
	name           string
	aliases        []string
	description    string
	stringResponse string
}

func SetCommands(s *discordgo.Session, m *discordgo.MessageCreate) {
	Info(s, m)
}

