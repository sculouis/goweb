package service

import (
	. "GoWeb/db"
	. "GoWeb/pojo"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get User
func FindAllUser(c *gin.Context) {
	var userList []User
	DB.Find(&userList)
	c.JSON(http.StatusOK, userList)
}

// Post User
func PostUser(c *gin.Context) {
	var user User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, "Error"+err.Error())
		return
	}
	DB.Create(&user)
	c.JSON(http.StatusOK, "successfully ok")
}

// delete User
func DeleteUser(c *gin.Context) {
	var user User
	if err := DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
