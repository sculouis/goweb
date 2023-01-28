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

func AddTodoRouter(r *gin.RouterGroup) {
	todo := r.Group("/todos")

	todo.GET("/", service.FindAllTodo)
	todo.POST("/", service.PostTodo)
	todo.DELETE("/:id", service.DeleteTodo)
}
