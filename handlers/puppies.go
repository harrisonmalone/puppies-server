package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/harrisonmalone/puppies-server/models"
)

type PuppyRoutes struct{}

func (r PuppyRoutes) Index(c *gin.Context) {
	var puppies []models.Puppy
	models.DB.Find(&puppies)
	c.JSON(http.StatusOK, puppies)
}

func (r PuppyRoutes) Create(c *gin.Context) {
	var body models.Puppy
	err := c.BindJSON(&body)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	models.DB.Create(&models.Puppy{
		Name: body.Name,
		Age:  body.Age,
		UserId: body.UserId,
	})
	c.Status(http.StatusCreated)
}

func (r PuppyRoutes) Show(c *gin.Context) {
	var puppy models.Puppy
	err := models.DB.First(&puppy, c.Param("id")).Error
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, puppy)
}

func (r PuppyRoutes) Destroy(c *gin.Context) {
	models.DB.Delete(&models.Puppy{}, c.Param("id"))
	c.Status(http.StatusNoContent)
}
