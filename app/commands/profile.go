package commands

import (
	"strconv"

	"github.com/blyndusk/cofy/app/core"
	"github.com/blyndusk/cofy/app/embeds"
	"github.com/blyndusk/cofy/app/services"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

type profileCommand struct {
	base           core.BaseCommand
	stringResponse string
}

func profile(s *discordgo.Session, m *discordgo.MessageCreate) {
	cmd := profileCommand{
		base: core.BaseCommand{
			Name:        "profile",
			Aliases:     []string{"profile", "pf"},
			Description: "Give informatiosn about you.",
		},
		stringResponse: "coffee",
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
	
	if strconv.Itoa(int(user.DiscordId)) != m.Author.ID {
		embeds.ProfileNotFound(s, m)
		services.PostUser(s, m)
	} else {
		embeds.Profile(s, m, user)
	}
}

func (cmd profileCommand) Help(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, cmd.base.Description)
}
