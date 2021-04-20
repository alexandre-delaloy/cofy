package middlewares

import (
	"net/http"

	"github.com/blyndusk/cofy/api/database"
	"github.com/blyndusk/cofy/api/helpers"
	"github.com/blyndusk/cofy/api/models"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context, input *models.UserInput) {
	if err := c.ShouldBindJSON(&input); err != nil {
		httpStatus, response := helpers.ErrorToJson(http.StatusBadRequest, err.Error())
		c.JSON(httpStatus, response)
		return
	}

	user := hydrateUser(input)

	if error := database.Db.Create(&user).Error; error != nil {
		httpStatus, response := helpers.GormErrorResponse(error)
		c.JSON(httpStatus, response)
		return
	}
}

func hydrateUser(input *models.UserInput) models.User {
	return models.User{
		Name:  input.Name,
		Coins: input.Coins,
		Xp:    input.Xp,
	}
}

func GetAllUsers(c *gin.Context, users *[]models.User) {
	if error := database.Db.Find(&users).Error; error != nil {
		httpStatus, response := helpers.GormErrorResponse(error)
		c.JSON(httpStatus, response)
		return
	}
}

func GetUserById(c *gin.Context, user *models.User) {

	if error := database.Db.First(&user, c.Params.ByName("id")).Error; error != nil {
		httpStatus, response := helpers.GormErrorResponse(error)
		c.JSON(httpStatus, response)
		return
	}
}

func DeleteUser(c *gin.Context, user *models.User) {

	if error := database.Db.First(&user, c.Params.ByName("id")).Delete(&user).Error; error != nil {
		httpStatus, response := helpers.GormErrorResponse(error)
		c.JSON(httpStatus, response)
		return
	}
}
