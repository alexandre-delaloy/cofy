package router

import (
	"github.com/blyndusk/cofy/api/controllers"
	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine) {
	usersRoute(r)
}

func usersRoute(r *gin.Engine) {
	r.POST("/users", controllers.CreateUser)

	r.GET("/users", controllers.GetAllUsers)
	r.GET("/users/:discord_id", controllers.GetUserById)

	r.PUT("/users/:discord_id", controllers.UpdateUser)

	r.DELETE("/users/:discord_id", controllers.DeleteUser)
}
