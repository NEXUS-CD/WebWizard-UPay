package controllers

import (
	"net/http"

	"UPay/services"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := services.FindUserByID(id)
	if err != nil {
		c.JSON(401, gin.H{"message": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}
