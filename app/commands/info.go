package commands

import (
	"github.com/blyndusk/cofy/app/core"
	"github.com/blyndusk/cofy/app/helpers"
	"github.com/blyndusk/cofy/app/services"

	"github.com/bwmarrin/discordgo"
)

type infoCommand struct {
	base           core.BaseCommand
	stringResponse string
}

func InfoCommandHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	cmd := infoCommand{
		base: core.BaseCommand{
			Name:        "info",
			Aliases:     []string{"info", "i"},
			Description: "Give informatiosn about Cofy bot.",
		},
	}

	if services.MatchExecuteCommand(s, m, cmd.base.Aliases) == true {
		services.CallExecute(cmd, s, m)
	}
	if services.MatchHelpCommand(s, m, cmd.base.Aliases) == true {
		services.CallHelp(cmd, s, m)
	}
}

func (cmd infoCommand) Execute(s *discordgo.Session, m *discordgo.MessageCreate) {
	helpers.EmbedViewInfos(s, m)
}

func (cmd infoCommand) Help(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, cmd.base.Description)
}
