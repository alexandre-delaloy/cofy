package services

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/blyndusk/cofy/app/helpers"
	"github.com/blyndusk/cofy/app/middlewares"
	"github.com/bwmarrin/discordgo"
)

func ViewTaggedUserProfile(s *discordgo.Session, m *discordgo.MessageCreate, aliases []string) {
	// get tagged user id from the command
	taggedUserId := GetUserIdFromCommand(m)
	// try to get the user with the tagged user id
	user := middlewares.GetUser(s, taggedUserId)
	// look for the discord user in discord session
	discordUser, _ := s.User(taggedUserId)
	// if both ids match, view the user
	if strconv.Itoa(user.DiscordId) == discordUser.ID {
		helpers.EmbedViewProfile(s, m, user, discordUser)
		// else,  view the user not found
	} else {
		helpers.EmbedProfileNotFound(s, m, user, discordUser)
	}
}

func GetUserIdFromCommand(m *discordgo.MessageCreate) string {
	res := strings.Split(m.Content, " ")
	if len(res) > 2 {
		taggedUser := res[2]
		isUserId, _ := regexp.Match(`<@![0-9]+>`, []byte(res[2]))
		if isUserId {
			userId := taggedUser
			userId = strings.Replace(userId, "<@!", "", 1)
			userId = strings.Replace(userId, ">", "", 1)
			return userId
		}
	}
	return ""
}
