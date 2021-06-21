package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/harrisonmalone/puppies-server/models"
)

type UserRoutes struct{}

func (r UserRoutes) Create(c *gin.Context) {
	var body models.User
	err := c.BindJSON(&body)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	models.DB.Create(&models.User{
		Email: body.Email,
		Password:  body.Password,
	})
	c.Status(http.StatusCreated)
}