package router

import (
	"github.com/blyndusk/cofy/api/controllers"
	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine) {
	usersRoute(r)
}

func usersRoute(r *gin.Engine) {
	r.GET("/users", controllers.GetAllUsers)
	r.GET("/users/:id", controllers.GetUserById)

	r.POST("/users", controllers.CreateUser)
}
