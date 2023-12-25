package main

import (
	"hospital/configs"
	"hospital/models"
)

func init() {
	configs.ConnectToDB()
}

func main() {
	configs.DB.AutoMigrate(&models.Person{}, &models.Patient{})
}
