package services

import (
	"fmt"
	"strconv"

	"github.com/blyndusk/cofy/app/core"
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

func EmbedViewProfile(s *discordgo.Session, m *discordgo.MessageCreate, user core.User, discordUser *discordgo.User) {
	logrus.Info("discord - user")
	logrus.Info(discordUser)
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Color:       1752220,
		Title:       fmt.Sprintf("%s's Cofy profile", discordUser.Username),
		Description: fmt.Sprintf("ID: %s", discordUser.ID),
		Thumbnail:   &discordgo.MessageEmbedThumbnail{URL: discordUser.AvatarURL(""), Width: 10, Height: 10},
		Fields: []*discordgo.MessageEmbedField{
			{Name: ":coin: CF:", Value: strconv.Itoa(user.Coins), Inline: true},
			{Name: ":chart_with_upwards_trend: XP:", Value: strconv.Itoa(user.Xp), Inline: true},
		},
		Footer: &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("Member since %s", user.CreatedAt.Format("01-02-2006"))},
	})

}

func EmbedCreatingProfile(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Color:       3447003,
		Title:       "Creating profile...",
	})

}

func EmbedViewGains(s *discordgo.Session, m *discordgo.MessageCreate, gainedCoins int, gainedXp int) {

	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Color: 1146986,
		Title: fmt.Sprintf(":coin: CF: +%d / :chart_with_upwards_trend: +%d", gainedCoins, gainedXp),
	})
}
