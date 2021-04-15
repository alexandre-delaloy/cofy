package commands

import (
	"fmt"

	"github.com/blyndusk/cofy/app/core"
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
		services.CallExecute(cmd, s, m)
	}
	if services.MatchCommand(true, s, m, cmd.base.Aliases) == true {
		services.CallHelp(cmd, s, m)
	}
}

func (command profileCommand) Execute(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, err := s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("%s's Cofy profile", m.Author.Username),
		Description: "desc desc",
		Thumbnail:   &discordgo.MessageEmbedThumbnail{URL: m.Author.AvatarURL(""), Width: 10, Height: 10},
		Footer:      &discordgo.MessageEmbedFooter{Text: "Source code: https://github.com/blyndusk/cofy"},
		Fields: []*discordgo.MessageEmbedField{
			{Name: ":coin: CF:", Value: "0", Inline: true},
			{Name: ":chart_with_upwards_trend: XP:", Value: "0", Inline: true},
			// {Name: "bar2", Value: "foo2", Inline: true},
		},
	})
	if err != nil {
		log.Error(err)
	}
}

func (cmd profileCommand) Help(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, cmd.base.Description)
}
