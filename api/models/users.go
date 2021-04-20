package models

import "time"

type User struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	Coins     uint      `json:"coins"`
	Xp        uint      `json:"xp"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Users []User
