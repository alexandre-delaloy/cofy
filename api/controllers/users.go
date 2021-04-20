package controllers

import (
	"github.com/blyndusk/cofy/api/database"
	"github.com/blyndusk/cofy/api/helpers"
	"github.com/blyndusk/cofy/api/models"
	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	var users []models.User

	if error := database.Db.Find(&users).Error; error != nil {
		httpStatus, response := helpers.GormErrorResponse(error)
		c.JSON(httpStatus, response)
		return
	}
}

func GetUserById(c *gin.Context) {
	var user models.User

	if error := database.Db.First(&user, c.Params.ByName("id")).Error; error != nil {
		httpStatus, response := helpers.GormErrorResponse(error)
		c.JSON(httpStatus, response)
		return
	}
}
