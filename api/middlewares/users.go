package middlewares

import (
	"net/http"

	"github.com/blyndusk/cofy/api/database"
	"github.com/blyndusk/cofy/api/helpers"
	"github.com/blyndusk/cofy/api/models"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func CreateUser(c *gin.Context, input *models.UserInput) {
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		httpStatus, response := helpers.ErrorToJson(http.StatusBadRequest, err.Error())
		c.JSON(httpStatus, response)
		return
	}

	user := hydrateUser(input)
	if err := database.Db.Create(&user).Error; err != nil {
		log.Error(err)
		httpStatus, response := helpers.GormErrorResponse(err)
		c.JSON(httpStatus, response)
		return
	}
}

func GetAllUsers(c *gin.Context, users *models.Users) {
	if err := database.Db.Find(&users).Error; err != nil {
		log.Error(err)
		httpStatus, response := helpers.GormErrorResponse(err)
		c.JSON(httpStatus, response)
		return
	}
}

func GetUserById(c *gin.Context, user *models.User) {
	if err := database.Db.Where("discord_id = ?", c.Params.ByName("discord_id")).First(&user).Error; err != nil {
		log.Error(err)
		httpStatus, response := helpers.GormErrorResponse(err)
		c.JSON(httpStatus, response)
		return
	}
}

func UpdateUser(c *gin.Context, user *models.User, input *models.UserInput) {
	GetUserById(c, user)
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		httpStatus, response := helpers.ErrorToJson(http.StatusBadRequest, err.Error())
		c.JSON(httpStatus, response)
		return
	}

	updatedUser := hydrateUser(input)
	database.Db.Model(&user).Updates(updatedUser)
}

func DeleteUser(c *gin.Context, user *models.User) {
	if err := database.Db.First(&user, c.Params.ByName("id")).Delete(&user).Error; err != nil {
		log.Error(err)
		httpStatus, response := helpers.GormErrorResponse(err)
		c.JSON(httpStatus, response)
		return
	}
}

func hydrateUser(input *models.UserInput) models.User {
	return models.User{
		DiscordId: input.DiscordId,
		Name:      input.Name,
		Coins:     input.Coins,
		Xp:        input.Xp,
		Level:     input.Level,
	}
}
