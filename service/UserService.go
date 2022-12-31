package service

import (
	"GoWeb/pojo"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var userList = []pojo.User{}

// Get User
func FindAllUser(c *gin.Context) {
	c.JSON(http.StatusOK, userList)
}

// Post User
func PostUser(c *gin.Context) {
	user := pojo.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, "Error"+err.Error())
		return
	}
	userList = append(userList, user)
	c.JSON(http.StatusOK, "successfully ok")
}

// delete User
func DeleteUser(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	for _, user := range userList {
		log.Println(user)
		if user.Id == userId {
			userList = append(userList[:userId], userList[userId+1:]...)
			c.JSON(http.StatusOK, "sucessfully deleted")
			return
		}
	}
	c.JSON(http.StatusNotFound, "fail to deleted")
}
