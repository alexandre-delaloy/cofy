package commands

import (
	"github.com/blyndusk/cofy/app/core"
	"github.com/blyndusk/cofy/app/services"
	log "github.com/sirupsen/logrus"

	"github.com/bwmarrin/discordgo"
)

type infoCommand struct {
	base           core.BaseCommand
	stringResponse string
}

func info(s *discordgo.Session, m *discordgo.MessageCreate) {
	cmd := infoCommand{
		base: core.BaseCommand{
			Name:        "info",
			Aliases:     []string{"info", "i"},
			Description: "Give informatiosn about Cofy bot.",
		},
		stringResponse: "Hi! I'm Cofy, a Discord bot made with the Golang programming language.\nMy purpose is to help you to gain cOoOOooOoooffEeEEeeEeee...",
	}

	if services.MatchCommand(false, s, m, cmd.base.Aliases) == true {
		log.Info("cmd | execute | info")
		services.CallExecute(cmd, s, m)
	}
	if services.MatchCommand(true, s, m, cmd.base.Aliases) == true {
		log.Info("cmd | help | info")
		services.CallHelp(cmd, s, m)
	}
}

func (cmd infoCommand) Execute(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, cmd.stringResponse)
}

func (cmd infoCommand) Help(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, cmd.base.Description)
}
