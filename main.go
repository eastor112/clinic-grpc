package main

import (
	"net/http"

	"github.com/eastor112/clinic-grpc/configs"
	"github.com/eastor112/clinic-grpc/middlewares"
	"github.com/eastor112/clinic-grpc/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	configs.ConnectToDB()
}

func main() {

	r := gin.Default()
	r.Use(middlewares.SecretMiddleware())

	routes.PersonRouter(r)
	routes.PatientRouter(r)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world from server Go.",
		})
	})
	r.Run()
}
