package service

import (
	. "GoWeb/db"
	. "GoWeb/pojo"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get All Todo
func FindAllTodo(c *gin.Context) {
	var todoList []Todo
	DB.Find(&todoList)
	c.JSON(http.StatusOK, todoList)
}

// Post Todo
func PostTodo(c *gin.Context) {
	var todo Todo
	err := c.BindJSON(&todo)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, "Error"+err.Error())
		return
	}
	DB.Create(&todo)
	c.JSON(http.StatusOK, "successfully ok")
}

// delete Todo
func DeleteTodo(c *gin.Context) {
	var todo Todo
	if err := DB.Where("id = ?", c.Param("id")).First(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	DB.Delete(&todo)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
