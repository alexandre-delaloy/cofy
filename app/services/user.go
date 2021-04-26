package services

import (
	"strconv"

	"github.com/blyndusk/cofy/app/core"
	"github.com/blyndusk/cofy/app/helpers"
	"github.com/blyndusk/cofy/app/middlewares"
	"github.com/bwmarrin/discordgo"
)

func HandleUserNotFound(s *discordgo.Session, m *discordgo.MessageCreate, user core.User) {
	// if the id of the user is not the author id, user doesnt exist
	if strconv.Itoa(user.DiscordId) != m.Author.ID {
		helpers.EmbedCreatingProfile(s, m)
		// creating user
		middlewares.CreateUser(m.Author)
		// get this new user
		newUser := middlewares.GetUser(s, m.Author.ID)
		// get the discord user from the session using discordId
		discordUser, _ := s.User(strconv.Itoa(newUser.DiscordId))
		// display profile
		helpers.EmbedViewProfile(s, m, newUser, discordUser)
	}
}
