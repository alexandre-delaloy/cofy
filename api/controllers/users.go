package controllers

import (
	"net/http"

	"github.com/blyndusk/cofy/api/middlewares"
	"github.com/blyndusk/cofy/api/models"
	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	var users []models.User

	middlewares.GetAllUsers(c, &users)
	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

func GetUserById(c *gin.Context) {
	var user models.User

	middlewares.GetUserById(c, &user)

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}
