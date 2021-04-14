package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func UseCommand(s *discordgo.Session, m *discordgo.MessageCreate, cmd Command) {
	var aliases = cmd.aliases
	if matchCommand(false, s, m, aliases) == true {
		s.ChannelMessageSend(m.ChannelID, cmd.stringResponse)
	}
	if matchCommand(true, s, m, aliases) == true {
		s.ChannelMessageSend(m.ChannelID, cmd.description)
	}
}

func matchCommand(isHelp bool, s *discordgo.Session, m *discordgo.MessageCreate, aliases []string) bool {
	if isHelp {
		for _, el := range aliases {
			if m.Content == fmt.Sprintf("%s help %s", Prefix, el) {
				return true
			}
		}
		return false
	} else {
		for _, el := range aliases {
			if m.Content == fmt.Sprintf("%s %s", Prefix, el) {
				return true
			}
		}
		return false
	}
}