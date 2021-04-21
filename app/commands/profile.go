package commands

import (
	"fmt"
	"strconv"

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
	if user.DiscordId != m.Author.ID {
		s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Title: "user not found; creating...",
		})
		services.PostUser(m.Author)
	} else {

		log.Info(m.Author.ID)
		s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Title:       fmt.Sprintf("%s's Cofy profile", m.Author.Username),
			Description: fmt.Sprintf("ID: %s", m.Author.ID),
			Thumbnail:   &discordgo.MessageEmbedThumbnail{URL: m.Author.AvatarURL(""), Width: 10, Height: 10},
			Fields: []*discordgo.MessageEmbedField{
				{Name: ":coin: CF:", Value: strconv.Itoa(user.Coins), Inline: true},
				{Name: ":chart_with_upwards_trend: XP:", Value: strconv.Itoa(user.Xp), Inline: true},
			},
			Footer: &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("Member since %s", user.CreatedAt.Format("01-02-2006"))},
		})
	}
}

func (cmd profileCommand) Help(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, cmd.base.Description)
}
