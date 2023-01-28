package src

import (
	"GoWeb/service"

	"github.com/gin-gonic/gin"
)

func AddUserRouter(r *gin.RouterGroup) {
	user := r.Group("/users")

	user.GET("/", service.FindAllUser)
	user.POST("/", service.PostUser)
	user.DELETE("/:id", service.DeleteUser)
}
