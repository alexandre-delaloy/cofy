package commands

import (
	"fmt"

	"github.com/blyndusk/cofy/app/core"
	"github.com/blyndusk/cofy/app/helpers"
	"github.com/blyndusk/cofy/app/middlewares"
	"github.com/blyndusk/cofy/app/services"

	"github.com/bwmarrin/discordgo"
)

type drinksCommand struct {
	base core.BaseCommand
}

func DrinksCommandHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	cmd := drinksCommand{
		base: core.BaseCommand{
			Name:        "drinks",
			Aliases:     []string{"drinks", "d"},
			Description: "Display drinks list",
		},
	}

	if services.MatchExecuteCommand(s, m, cmd.base.Aliases) == true {
		services.CallExecute(cmd, s, m)
	}
	if services.MatchHelpCommand(s, m, cmd.base.Aliases) == true {
		services.CallHelp(cmd, s, m)
	}
}

func (cmd drinksCommand) Execute(s *discordgo.Session, m *discordgo.MessageCreate) {
	// get all drinks
	drinks := middlewares.GetDrinks(s)
	// decalre empty array
	var fields []*discordgo.MessageEmbedField
	// for every drink
	for _, s := range drinks {
		// add discord embed field in array
		fields = append(fields, &discordgo.MessageEmbedField{
			Name:  fmt.Sprintf("%s %s", s.Emoji, s.Name),
			Value: fmt.Sprintf(":coin: : %d - :trophy: : %d", s.Price, s.RequiredLevel),
		})
	}
	// display drinks
	helpers.EmbedViewDrinks(s, m, fields)
}

func (cmd drinksCommand) Help(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, cmd.base.Description)
}
