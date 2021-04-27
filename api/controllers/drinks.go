package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/blyndusk/cofy/api/middlewares"
	"github.com/blyndusk/cofy/api/models"
)

func CreateDrink(c *gin.Context) {
	var input models.DrinkInput
	middlewares.CreateDrink(c, &input)
	c.JSON(http.StatusOK, input)
}

func GetAllDrinks(c *gin.Context) {
	var drinks models.Drinks
	middlewares.GetAllDrinks(c, &drinks)
	c.JSON(http.StatusOK, drinks)
}

func GetDrinkById(c *gin.Context) {
	var drink models.Drink
	middlewares.GetDrinkById(c, &drink)
	c.JSON(http.StatusOK, drink)
}

func UpdateDrink(c *gin.Context) {
	var drink models.Drink
	var input models.DrinkInput
	middlewares.UpdateDrink(c, &drink, &input)
	c.JSON(http.StatusOK, drink)
}

func DeleteDrink(c *gin.Context) {
	var drink models.Drink
	middlewares.DeleteDrink(c, &drink)
	c.JSON(http.StatusOK, gin.H{
		"message": "Drink deleted successfully",
	})
}
