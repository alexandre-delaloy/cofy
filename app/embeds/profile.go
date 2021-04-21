package embeds

import (
	"fmt"
	"strconv"

	"github.com/blyndusk/cofy/app/core"
	"github.com/bwmarrin/discordgo"
)

func ProfileNotFound(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Color:       15158332,
		Title:       "Profile not found !",
		Description: "Looks like you don't have any profile.",
	})
}

func ProfileCreating(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Color:       15105570,
		Title:       "Creating profile...",
		Description: "We're creating your new profile.",
	})
}

func ProfileCreated(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Color:       3066993,
		Title:       "Profile created !",
		Description: "You're profile has been successfully created. Enjoy !",
	})
}

func Profile(s *discordgo.Session, m *discordgo.MessageCreate, user core.User) {

	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Color: 1752220,
		Title:       fmt.Sprintf("%s's Cofy profile", m.Author.Username),
		Description: fmt.Sprintf("ID: %s", m.Author.ID),
		Thumbnail:   &discordgo.MessageEmbedThumbnail{URL: m.Author.AvatarURL(""), Width: 10, Height: 10},
		Fields: []*discordgo.MessageEmbedField{
			{Name: ":coin: CF:", Value: strconv.Itoa(user.Coins), Inline: true},
			{Name: ":chart_with_upwards_trend: XP:", Value: strconv.Itoa(user.Xp), Inline: true},
		},
		Footer: &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("Member since %s", user.CreatedAt.Format("01-02-2006"))},
	})

}
