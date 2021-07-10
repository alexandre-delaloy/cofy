package models

type Drink struct {
	Id            uint   `json:"id"`
	Name          string `json:"name"`
	Emoji         string `json:"emoji"`
	Price         uint   `json:"price"`
	RequiredLevel uint   `json:"required_level"`
}

type Drinks []Drink

type DrinkInput struct {
	Name          string `json:"name"`
	Emoji         string `json:"emoji"`
	Price         uint   `json:"price"`
	RequiredLevel uint   `json:"required_level"`
}
