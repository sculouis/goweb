package service

import (
	"GoWeb/pojo"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get User
func FindAllUser(c *gin.Context) {
	var userList []pojo.User
	pojo.DB.Find(&userList)
	c.JSON(http.StatusOK, userList)
}

// Post User
func PostUser(c *gin.Context) {
	var user pojo.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, "Error"+err.Error())
		return
	}
	pojo.DB.Create(&user)
	c.JSON(http.StatusOK, "successfully ok")
}

// delete User
func DeleteUser(c *gin.Context) {
	var user pojo.User
	if err := pojo.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	pojo.DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
