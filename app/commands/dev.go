package commands

import (
	"fmt"
	"strconv"

	"github.com/blyndusk/cofy/app/core"
	"github.com/blyndusk/cofy/app/services"
	log "github.com/sirupsen/logrus"

	"github.com/bwmarrin/discordgo"
)

type devCommand struct {
	base           core.BaseCommand
	stringResponse string
}

func dev(s *discordgo.Session, m *discordgo.MessageCreate) {
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
	user := services.GetUser(m.Author.ID)

	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("%s's Cofy profile", m.Author.Username),
		Description: "desc desc",
		Thumbnail:   &discordgo.MessageEmbedThumbnail{URL: m.Author.AvatarURL(""), Width: 10, Height: 10},
		// Footer:      &discordgo.MessageEmbedFooter{Text: "Source code: https://github.com/blyndusk/cofy"},
		Fields: []*discordgo.MessageEmbedField{
			{Name: ":coin: CF:", Value: strconv.Itoa(user.Coins), Inline: true},
			{Name: ":chart_with_upwards_trend: XP:", Value: strconv.Itoa(user.Xp), Inline: true},
		},
	})

}

func (cmd devCommand) Help(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, cmd.base.Description)
}
