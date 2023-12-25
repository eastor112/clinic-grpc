package main

import (
	"github.com/eastor112/clinic-grpc/configs"
	"github.com/eastor112/clinic-grpc/models"
)

func init() {
	configs.ConnectToDB()
}

func main() {
	configs.DB.AutoMigrate(&models.Person{}, &models.Patient{})
}
