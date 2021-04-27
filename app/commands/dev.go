package commands

import (
	"fmt"

	"github.com/blyndusk/cofy/app/core"
	"github.com/blyndusk/cofy/app/middlewares"
	"github.com/blyndusk/cofy/app/services"
	"github.com/sirupsen/logrus"

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
	drinks := 	middlewares.GetDrinks(s)
	// logrus.Info(drinks)
	var fields []*discordgo.MessageEmbedField
	for _, s := range drinks {
		field := discordgo.MessageEmbedField{
			Name: fmt.Sprintf("%s %s", s.Emoji,  s.Name),
			Value: fmt.Sprintf(":coin: : %d - :trophy: : %d", s.Price,  s.RequiredLevel),
		}
		logrus.Info(field)
		fields = append(fields, &field)
	}
	logrus.Info(fields)
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Color:       1752220,
		Title:      "Drinks",
		// Description: discordUser.Discriminator,
		// Thumbnail:   &discordgo.MessageEmbedThumbnail{URL: discordUser.AvatarURL(""), Width: 10, Height: 10},
		Fields: fields,
		// Footer: &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("Member since %s", user.CreatedAt.Format("01-02-2006"))},
	})

	// for i := 0; i < 11; i++ {
	// 	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%d - %f", i, helpers.GetXpFromLevel(i)))

		
	// }
}

func (cmd devCommand) Help(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, cmd.base.Description)
}
