package commands

import (
	"github.com/bwmarrin/discordgo"
)

func SetCommands(s *discordgo.Session, m *discordgo.MessageCreate) {
	info(s, m)
	profile(s, m)
	dev(s, m)
}
