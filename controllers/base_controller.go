package controllers

import (
	"net/http"

	"UPay/services"

	"github.com/gin-gonic/gin"
)

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := services.FindUserByID(id)
	if err != nil {
		c.JSON(

			401,
			gin.H{"message": "User not found"},
		)
		return
	}
	c.JSON(http.StatusOK, user)
}
