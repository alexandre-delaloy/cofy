package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)


func MatchCommand(isHelp bool, s *discordgo.Session, m *discordgo.MessageCreate, aliases []string) bool {
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