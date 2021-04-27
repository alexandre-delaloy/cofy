package router

import (
	"github.com/blyndusk/cofy/api/controllers"
	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine) {
	usersRoute(r)
	drinksRoute(r)
}

func usersRoute(r *gin.Engine) {
	r.POST("/users", controllers.CreateUser)

	r.GET("/users", controllers.GetAllUsers)
	r.GET("/users/:discord_id", controllers.GetUserById)

	r.PUT("/users/:discord_id", controllers.UpdateUser)

	r.DELETE("/users/:discord_id", controllers.DeleteUser)
}


func drinksRoute(r *gin.Engine) {
	r.POST("/drinks", controllers.CreateDrink)

	r.GET("/drinks", controllers.GetAllDrinks)
	r.GET("/drinks/:id", controllers.GetDrinkById)

	r.PUT("/drinks/:id", controllers.UpdateDrink)

	r.DELETE("/drinks/:id", controllers.DeleteDrink)
}
