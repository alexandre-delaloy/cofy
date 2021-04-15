package core

import "github.com/bwmarrin/discordgo"

var Prefix string = "c!"
type BaseCommand struct {
	Name           string
	Aliases        []string
	Description    string
}
type ICommand interface {
	Execute(s *discordgo.Session, m *discordgo.MessageCreate)
	Help(s *discordgo.Session, m *discordgo.MessageCreate)
}
