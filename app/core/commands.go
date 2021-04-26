package core

import "github.com/bwmarrin/discordgo"

// command prefix used in all app
var Prefix string = "c!"

// the command struct, with essential informations
type BaseCommand struct {
	Name        string
	Aliases     []string
	Description string
}

// the command interface, with custom methods
type ICommand interface {
	Execute(s *discordgo.Session, m *discordgo.MessageCreate)
	Help(s *discordgo.Session, m *discordgo.MessageCreate)
}
