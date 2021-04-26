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
		Description: fmt.Sprintf("Level: %d", user.Level),
		Thumbnail:   &discordgo.MessageEmbedThumbnail{URL: discordUser.AvatarURL(""), Width: 10, Height: 10},
		Fields: []*discordgo.MessageEmbedField{
			{Name: ":coin: CF:", Value: strconv.Itoa(user.Coins), Inline: true},
			{Name: ":chart_with_upwards_trend: XP:", Value: strconv.Itoa(user.Xp), Inline: true},
		},
		Footer: &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("Member since %s", user.CreatedAt.Format("01-02-2006"))},
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

func EmbedCreatingProfile(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Color: 3447003,
		Title: "Creating profile...",
	})
}

func EmbedViewGains(s *discordgo.Session, m *discordgo.MessageCreate, gains core.Gains) {
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Color: 1146986,
		Title: fmt.Sprintf("[DEBUG] :coin: CF: +%d / :chart_with_upwards_trend: +%d", gains.Coins, gains.Xp),
	})
}

func EmbedViewInfos(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Color:       1146986,
		Title:       "Hi, I'm Cofy, a bot boosted in caffeine",
		Description: "I'm a Discord bot made with the Golang programming language. I'm a side-project to improve my creator's developping skills.",
		Thumbnail:   &discordgo.MessageEmbedThumbnail{URL: s.State.User.AvatarURL(""), Width: 6, Height: 6},
		Fields: []*discordgo.MessageEmbedField{
			{Name: ":coin: CF:", Value: "Each time you send a message, you gain coins. Coins are used to buy coffee."},
			{Name: ":chart_with_upwards_trend: XP:", Value: "Each time you send a message, you gain XP. XP makes you level up. Levels are used to unlock new coffee."},
		},
		Footer: &discordgo.MessageEmbedFooter{Text: "You can check my source code at https://github.com/blyndusk/cofy"},
	})
}
