package commands

import (
	"fmt"

	"github.com/blyndusk/cofy/app/core"
	"github.com/blyndusk/cofy/app/helpers"
	"github.com/blyndusk/cofy/app/services"

	"github.com/bwmarrin/discordgo"
)

type devCommand struct {
	base           core.BaseCommand
	stringResponse string
}

func DevCommandHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	cmd := devCommand{
		base: core.BaseCommand{
			Name:        "dev",
			Aliases:     []string{"dev", "d", "!"},
			Description: "Developper stuff",
		},
		stringResponse: "API: http://localhost:3003\nADM: http://localhost:3002",
	}

	if services.MatchExecuteCommand(s, m, cmd.base.Aliases) == true {
		services.CallExecute(cmd, s, m)
	}
	if services.MatchHelpCommand(s, m, cmd.base.Aliases) == true {
		services.CallHelp(cmd, s, m)
	}
}

func (cmd devCommand) Execute(s *discordgo.Session, m *discordgo.MessageCreate) {
	// s.ChannelMessageSend(m.ChannelID, cmd.stringResponse)
	for i := 0; i < 11; i++ {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%d - %f", i, helpers.GetXpFromLevel(i)))

	}
}

func (cmd devCommand) Help(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, cmd.base.Description)
}
