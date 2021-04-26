package helpers

import (
	"fmt"
	"strconv"

	"github.com/blyndusk/cofy/app/core"
	"github.com/bwmarrin/discordgo"
)

func EmbedViewProfile(s *discordgo.Session, m *discordgo.MessageCreate, user core.User, discordUser *discordgo.User) {
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
		Color: 3447003,
		Title: "Creating profile...",
	})

}

func EmbedViewGains(s *discordgo.Session, m *discordgo.MessageCreate, gainedCoins int, gainedXp int) {

	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Color: 1146986,
		Title: fmt.Sprintf("[DEBUG] :coin: CF: +%d / :chart_with_upwards_trend: +%d", gainedCoins, gainedXp),
	})
}

func EmbedProfileNotFound(s *discordgo.Session, m *discordgo.MessageCreate, user core.User, discordUser *discordgo.User) {
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Color:       15158332,
		Title:       fmt.Sprintf("%s's Cofy profile not found", discordUser.Username),
		Description: "This user has not created a profile yet. ",
		Thumbnail:   &discordgo.MessageEmbedThumbnail{URL: discordUser.AvatarURL(""), Width: 10, Height: 10},
		Footer:      &discordgo.MessageEmbedFooter{Text: "For a user to create a profile, he just need to send a message and his profile will be created automatically."},
	})
}