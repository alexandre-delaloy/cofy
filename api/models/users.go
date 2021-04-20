package models

import "time"

type User struct {
	Id        uint      `json:"id"`
	DiscordId string    `json:"discord_id"`
	Name      string    `json:"name"`
	Coins     uint      `json:"coins"`
	Xp        uint      `json:"xp"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Users []User

type UserInput struct {
	DiscordId string `json:"discord_id"`
	Name      string `json:"name"`
	Coins     uint   `json:"coins"`
	Xp        uint   `json:"xp"`
}
