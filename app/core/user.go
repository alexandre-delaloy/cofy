package core

import "time"

type User struct {
	Id        uint      `json:"id"`
	DiscordId int       `json:"discord_id"`
	Name      string    `json:"name"`
	Coins     int       `json:"coins"`
	Xp        int       `json:"xp"`
	Level     int       `json:"level"`
	CreatedAt time.Time `json:"created_at"`
}

type Gains struct {
	Coins int `json:"coins"`
	Xp    int `json:"xp"`
	Level int `json:"level"`
}
