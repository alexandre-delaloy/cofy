package commands

import (
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
	services.SeeOtheruserProfile(s, m, cmd.base.Aliases)
}

func (command profileCommand) Execute(s *discordgo.Session, m *discordgo.MessageCreate) {
	user := services.GetUser(s, m.Author.ID)

	services.CreateUserIfDoesntExist(user, s, m)
	if user.DiscordId == m.Author.ID {
		discordUser, _ := s.User(user.DiscordId)
		services.EmbedViewProfile(s, m, user, discordUser)
	}

}

func (cmd profileCommand) Help(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, cmd.base.Description)
}
