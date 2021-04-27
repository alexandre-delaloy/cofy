package middlewares

import (
	"net/http"

	"github.com/blyndusk/cofy/api/database"
	"github.com/blyndusk/cofy/api/helpers"
	"github.com/blyndusk/cofy/api/models"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func CreateDrink(c *gin.Context, input *models.DrinkInput) {
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		httpStatus, response := helpers.ErrorToJson(http.StatusBadRequest, err.Error())
		c.JSON(httpStatus, response)
		return
	}

	drink := hydrateDrink(input)
	if err := database.Db.Create(&drink).Error; err != nil {
		log.Error(err)
		httpStatus, response := helpers.GormErrorResponse(err)
		c.JSON(httpStatus, response)
		return
	}
}

func GetAllDrinks(c *gin.Context, drinks *models.Drinks) {
	if err := database.Db.Find(&drinks).Error; err != nil {
		log.Error(err)
		httpStatus, response := helpers.GormErrorResponse(err)
		c.JSON(httpStatus, response)
		return
	}
}

func GetDrinkById(c *gin.Context, drink *models.Drink) {
	if err := database.Db.Where("id = ?", c.Params.ByName("id")).First(&drink).Error; err != nil {
		log.Error(err)
		httpStatus, response := helpers.GormErrorResponse(err)
		c.JSON(httpStatus, response)
		return
	}
}

func UpdateDrink(c *gin.Context, drink *models.Drink, input *models.DrinkInput) {
	GetDrinkById(c, drink)
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		httpStatus, response := helpers.ErrorToJson(http.StatusBadRequest, err.Error())
		c.JSON(httpStatus, response)
		return
	}

	updateDrink := hydrateDrink(input)
	database.Db.Model(&drink).Updates(updateDrink)
}

func DeleteDrink(c *gin.Context, drink *models.Drink) {
	if err := database.Db.First(&drink, c.Params.ByName("id")).Delete(&drink).Error; err != nil {
		log.Error(err)
		httpStatus, response := helpers.GormErrorResponse(err)
		c.JSON(httpStatus, response)
		return
	}
}

func hydrateDrink(input *models.DrinkInput) models.Drink {
	return models.Drink{
		Name:          input.Name,
		Emoji:         input.Emoji,
		Price:         input.Price,
		RequiredLevel: input.RequiredLevel,
	}
}
