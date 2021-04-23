package services

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/blyndusk/cofy/app/core"
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
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

func SeeOtheruserProfile(s *discordgo.Session, m *discordgo.MessageCreate, aliases []string) {
	otherUserId := GetIdOfOtherUser(m)
	logrus.Info(s.User(otherUserId))
	logrus.Info(otherUserId)
	doesMatch := false
	for _, el := range aliases {
		MatchCommand, _ := regexp.Match(fmt.Sprintf("%s %s <@![0-9]+>", core.Prefix, el), []byte(m.Content))
		if MatchCommand && otherUserId != "" {
			logrus.Info("oui")
			doesMatch = true
		}
	}
	logrus.Info(doesMatch)
	if doesMatch {
		user := GetUser(s, otherUserId)
		discordUser, _ := s.User(otherUserId)
		EmbedViewProfile(s, m, user, discordUser)
	}
}

func GetIdOfOtherUser(m *discordgo.MessageCreate) string {
	res := strings.Split(m.Content, " ")
	if len(res) > 2 {
		taggedUser := res[2]
		logrus.Info(taggedUser)
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
