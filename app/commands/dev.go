package commands

import (
	"github.com/blyndusk/cofy/app/core"
	"github.com/blyndusk/cofy/app/services"
	log "github.com/sirupsen/logrus"

	"github.com/bwmarrin/discordgo"
)

type devCommand struct {
	base           core.BaseCommand
	stringResponse string
}

func Dev(s *discordgo.Session, m *discordgo.MessageCreate) {
	cmd := devCommand{
		base: core.BaseCommand{
			Name:        "dev",
			Aliases:     []string{"dev", "d", "!"},
			Description: "Developper stuff",
		},
		stringResponse: "API: http://localhost:3003\nADM: http://localhost:3002",
	}

	if services.MatchCommand(false, s, m, cmd.base.Aliases) == true {
		log.Info("cmd | execute | dev")
		services.CallExecute(cmd, s, m)
	}
	if services.MatchCommand(true, s, m, cmd.base.Aliases) == true {
		log.Info("cmd | help | idevfo")
		services.CallHelp(cmd, s, m)
	}
}

func (cmd devCommand) Execute(s *discordgo.Session, m *discordgo.MessageCreate) {
	user := services.GetUser(s, m.Author.ID)
	// t1, _:= time.Parse(
	// 	time.RFC3339,
	// 	user.CreatedAt)
	log.Info(user.CreatedAt.Format("01-02-2006"))

}

func (cmd devCommand) Help(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, cmd.base.Description)
}
