package services

import (
	"fmt"

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
