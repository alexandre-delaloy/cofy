package services

import (
	"fmt"
	"regexp"

	"github.com/blyndusk/cofy/app/core"
	"github.com/bwmarrin/discordgo"
)

// bind execute method with command interface
func CallExecute(ICommand core.ICommand, s *discordgo.Session, m *discordgo.MessageCreate) {
	ICommand.Execute(s, m)
}

// bind help method with command interface
func CallHelp(ICommand core.ICommand, s *discordgo.Session, m *discordgo.MessageCreate) {
	ICommand.Help(s, m)
}

func MatchExecuteCommand(s *discordgo.Session, m *discordgo.MessageCreate, aliases []string) bool {
	// for every alias
	for _, el := range aliases {
		// e.g. `c! profile || c! pf`
		if m.Content == fmt.Sprintf("%s %s", core.Prefix, el) {
			return true
		}
	}
	return false
}
func MatchHelpCommand(s *discordgo.Session, m *discordgo.MessageCreate, aliases []string) bool {
	for _, el := range aliases {
		// e.g. `c! help profile || c! help pf`
		if m.Content == fmt.Sprintf("%s help %s", core.Prefix, el) {
			return true
		}
	}
	return false

}

func MatchTaggedUserProfileCommand(s *discordgo.Session, m *discordgo.MessageCreate, aliases []string) bool {
	for _, el := range aliases {
		// e.g. `c! profile <@!831597490851020820> `
		MatchCommand, _ := regexp.Match(fmt.Sprintf("%s %s <@![0-9]+>", core.Prefix, el), []byte(m.Content))
		if MatchCommand {
			return true
		}
	}
	return false
}
