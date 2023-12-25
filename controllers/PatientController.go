package controllers

import (
	"fmt"
	"hospital/configs"
	"hospital/models"

	"github.com/gin-gonic/gin"
)

type PatientRequestBody struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Phone     string `json:"phone"`
	SSN       string `json:"ssn"`
}

func PatientCreate(c *gin.Context) {

	body := PatientRequestBody{}

	c.BindJSON(&body)
	patient := &models.Patient{FirstName: body.FirstName, LastName: body.LastName, Phone: body.Phone}
	result := configs.DB.Create(&patient)

	if result.Error != nil {
		c.JSON(500, gin.H{"Error": "Failed to insert"})
		return
	}

	c.JSON(200, &patient)
}

func PatientGet(c *gin.Context) {
	hospital, exists1 := c.Get("Hospital")
	protection, exists2 := c.Get("Protection")
	fmt.Printf("%v %v \n", hospital, exists1)
	fmt.Printf("%v %v \n", protection, exists2)

	q := c.Query("q")

	fmt.Printf("%v \n", q)

	var patients []models.Patient
	configs.DB.Find(&patients)
	c.JSON(200, patients)
}

func PatientGetById(c *gin.Context) {
	id := c.Param("id")
	var patient models.Patient
	result := configs.DB.First(&patient, id)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Patient not found"})
		return
	}

	c.JSON(200, patient)
}

func PatientPatch(c *gin.Context) {
	id := c.Param("id")

	var patient models.Patient
	if err := configs.DB.First(&patient, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Patient not found"})
		return
	}

	var body PatientRequestBody
	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": "Json format is not valid"})
		return
	}

	if body.LastName != "" {
		patient.LastName = body.LastName
	}
	if body.FirstName != "" {
		patient.FirstName = body.FirstName
	}
	if body.Phone != "" {
		patient.Phone = body.Phone
	}

	if err := configs.DB.Save(&patient).Error; err != nil {
		c.JSON(500, gin.H{"error": "Something went wrong"})
		return
	}

	c.JSON(200, &patient)
}

func PatientDelete(c *gin.Context) {
	id := c.Param("id")
	var patient models.Patient
	result := configs.DB.Delete(&patient, id)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Something went wrong"})
		return
	}
	c.JSON(200, gin.H{"deleted": true})
}
