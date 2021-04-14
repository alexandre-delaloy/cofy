package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

func Profile(s *discordgo.Session, m *discordgo.MessageCreate, prefix string) {
	command := "pf"
	description := "Give informations about you"
	response := "info"

	if m.Content == fmt.Sprintf("%s %s", prefix, command) {
		// Foo := 
		
		s.ChannelMessageSend(m.ChannelID, response)
		log.Info(m.Author.AvatarURL)
		_, err := s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Title:       fmt.Sprintf("%s's Cofy profile", m.Author.Username),
			Description: "desc desc",
			Thumbnail: &discordgo.MessageEmbedThumbnail{URL: 	m.Author.AvatarURL(""), Width: 10, Height: 10},
			Footer:      &discordgo.MessageEmbedFooter{Text: "foo"},
			Fields:      []*discordgo.MessageEmbedField{
				{Name: ":coin: CF:", Value: "0", Inline: true},
				{Name: ":chart_with_upwards_trend: XP:", Value: "0", Inline: true},
				// {Name: "bar2", Value: "foo2", Inline: true},
			},
		})
		if err != nil {
			logrus.Error(err)
		}
	}

	if m.Content == fmt.Sprintf("%s help %s", prefix, command) {
		s.ChannelMessageSend(m.ChannelID, description)
	}
}
