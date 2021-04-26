package services

import (
	"fmt"
	"regexp"

	"github.com/blyndusk/cofy/app/core"
	"github.com/bwmarrin/discordgo"
)

func CallExecute(ICommand core.ICommand, s *discordgo.Session, m *discordgo.MessageCreate) {
	ICommand.Execute(s, m)
}

func CallHelp(ICommand core.ICommand, s *discordgo.Session, m *discordgo.MessageCreate) {
	ICommand.Help(s, m)
}

func MatchCommand(isHelp bool, s *discordgo.Session, m *discordgo.MessageCreate, aliases []string) bool {
	if isHelp {
		for _, el := range aliases {
			if m.Content == fmt.Sprintf("%s help %s", core.Prefix, el) {
				return true
			}
		}
		return false
	} else {
		for _, el := range aliases {
			if m.Content == fmt.Sprintf("%s %s", core.Prefix, el) {
				return true
			}
		}
		return false
	}
}

func MatchTaggedUserProfileCommand(s *discordgo.Session, m *discordgo.MessageCreate, aliases []string) bool {
	for _, el := range aliases {
		MatchCommand, _ := regexp.Match(fmt.Sprintf("%s %s <@![0-9]+>", core.Prefix, el), []byte(m.Content))
		if MatchCommand {
			return true
		}
	}
	return false
}
