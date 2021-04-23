package core

import "time"

type User struct {
	Id        uint      `json:"id"`
	DiscordId string    `json:"discord_id"`
	Name      string    `json:"name"`
	Coins     int       `json:"coins"`
	Xp        int       `json:"xp"`
	CreatedAt time.Time `json:"created_at"`
}
