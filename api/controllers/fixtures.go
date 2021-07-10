package controllers

import (
	"net/http"

	"github.com/blyndusk/cofy/api/database"
	"github.com/blyndusk/cofy/api/models"
	"github.com/gin-gonic/gin"
)

func LoadData(c *gin.Context) {
	users := models.Users{
		{
			Name:      "Cofy",
			DiscordId: 831597490851020820,
			Coins:     666666,
			Xp:        3666,
			Level:     10,
		},
	}

	drinks := models.Drinks{
		{Name: "Coffee", Emoji: ":coffee:", Price: 10, RequiredLevel: 1},
		{Name: "Tea", Emoji: ":tea:", Price: 20, RequiredLevel: 2},
		{Name: "Beer", Emoji: ":beer:", Price: 50, RequiredLevel: 5},
		{Name: "Wine", Emoji: ":wine_glass:", Price: 80, RequiredLevel: 8},
		{Name: "Whisky", Emoji: ":tumbler_glass:", Price: 100, RequiredLevel: 10},
	}
	database.Db.Create(&users)
	database.Db.Create(&drinks)
	c.JSON(http.StatusOK, gin.H{
		"message": "Fixtures loaded",
	})

}
