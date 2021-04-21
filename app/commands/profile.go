package commands

import (
	"strconv"

	"github.com/blyndusk/cofy/app/core"
	"github.com/blyndusk/cofy/app/services"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

type profileCommand struct {
	base core.BaseCommand
}

func Profile(s *discordgo.Session, m *discordgo.MessageCreate) {
	cmd := profileCommand{
		base: core.BaseCommand{
			Name:        "profile",
			Aliases:     []string{"profile", "pf"},
			Description: "Give informations about you.",
		},
	}

	if services.MatchCommand(false, s, m, cmd.base.Aliases) == true {
		log.Info("cmd | execute | profile")
		services.CallExecute(cmd, s, m)
	}
	if services.MatchCommand(true, s, m, cmd.base.Aliases) == true {
		log.Info("cmd | help | profile")
		services.CallHelp(cmd, s, m)
	}
}

func (command profileCommand) Execute(s *discordgo.Session, m *discordgo.MessageCreate) {
	user := services.GetUser(m.Author.ID)

	services.CreateUserIfDoesntExist(user, s, m)
	if strconv.Itoa(int(user.DiscordId)) == m.Author.ID {
		services.EmbedViewProfile(s, m, user)
	}
}

func (cmd profileCommand) Help(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, cmd.base.Description)
}
