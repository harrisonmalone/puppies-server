package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/harrisonmalone/puppies-server/models"
)

type PuppiesRoutes struct{}

func (r PuppiesRoutes) Index(c *gin.Context) {
	var puppies []models.Puppy
	models.DB.Find(&puppies)
	c.JSON(http.StatusOK, puppies)
}

func (r PuppiesRoutes) Create(c *gin.Context) {
	var body models.Puppy
	err := c.BindJSON(&body)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	models.DB.Create(&models.Puppy{
		Name: body.Name,
		Age:  body.Age,
	})
	c.Status(http.StatusCreated)
}

func (r PuppiesRoutes) Show(c *gin.Context) {
	var puppy models.Puppy
	err := models.DB.First(&puppy, c.Param("id")).Error
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, puppy)
}

func (r PuppiesRoutes) Destroy(c *gin.Context) {
	models.DB.Delete(&models.Puppy{}, c.Param("id"))
	c.Status(http.StatusNoContent)
}
