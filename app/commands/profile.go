package commands

import (
	"strconv"

	"github.com/blyndusk/cofy/app/core"
	"github.com/blyndusk/cofy/app/helpers"
	"github.com/blyndusk/cofy/app/middlewares"
	"github.com/blyndusk/cofy/app/services"
	"github.com/bwmarrin/discordgo"
)

type profileCommand struct {
	base core.BaseCommand
}

func ProfileCommandHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// init new command
	cmd := profileCommand{
		base: core.BaseCommand{
			Name:        "profile",
			Aliases:     []string{"profile", "pf"},
			Description: "Give informations about you.",
		},
	}
	// core methods
	if services.MatchExecuteCommand(s, m, cmd.base.Aliases) == true {
		services.CallExecute(cmd, s, m)
	}

	if services.MatchHelpCommand(s, m, cmd.base.Aliases) == true {
		services.CallHelp(cmd, s, m)
	}

	if services.MatchTaggedUserProfileCommand(s, m, cmd.base.Aliases) == true {
		services.ViewTaggedUserProfile(s, m, cmd.base.Aliases)
	}
}

// execute method
func (cmd profileCommand) Execute(s *discordgo.Session, m *discordgo.MessageCreate) {
	// systematically get user each time someone createw a message
	user := middlewares.GetUser(s, m.Author.ID)
	// if the id of the user is not the author id, user doesnt exist
	if strconv.Itoa(user.DiscordId) != m.Author.ID {
		// if the user is not found, create a new user in db
		services.HandleUserNotFound(s, m, user)
	} else {
		// get the discord user from the session using discordId
		discordUser, _ := s.User(strconv.Itoa(user.DiscordId))
		// display profile
		helpers.EmbedViewProfile(s, m, user, discordUser)
	}
}

// help method
func (cmd profileCommand) Help(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, cmd.base.Description)
}
