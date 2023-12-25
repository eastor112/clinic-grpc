package controllers

import (
	"github.com/eastor112/clinic-grpc/configs"
	"github.com/eastor112/clinic-grpc/models"

	"github.com/gin-gonic/gin"
)

type PersonRequestBody struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

func PersonCreate(c *gin.Context) {

	body := PersonRequestBody{}

	c.BindJSON(&body)
	person := &models.Person{Name: body.Name, Address: body.Address, Phone: body.Phone}
	result := configs.DB.Create(&person)

	if result.Error != nil {
		c.JSON(500, gin.H{"Error": "Failed to insert"})
		return
	}

	c.JSON(200, &person)
}

func PersonGet(c *gin.Context) {
	var persons []models.Person
	configs.DB.Find(&persons)
	c.JSON(200, &persons)
}

func PersonGetById(c *gin.Context) {
	id := c.Param("id")
	var person models.Person
	configs.DB.First(&person, id)
	c.JSON(200, &person)
}

func PersonUpdate(c *gin.Context) {

	id := c.Param("id")
	var person models.Person
	configs.DB.First(&person, id)

	body := PersonRequestBody{}
	c.BindJSON(&body)
	data := &models.Person{Name: body.Name, Address: body.Address, Phone: body.Phone}

	result := configs.DB.Model(&person).Updates(data)

	if result.Error != nil {
		c.JSON(500, gin.H{"Error": true, "message": "Failed to update"})
		return
	}
	c.JSON(200, &person)
}

func PersonDelete(c *gin.Context) {
	id := c.Param("id")
	var person models.Person
	result := configs.DB.Delete(&person, id)

	if result.Error != nil {
		c.JSON(500, gin.H{"Error": true, "message": "Failed to delete"})
		return
	}
	c.JSON(200, gin.H{"deleted": true})
}
