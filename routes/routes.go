package routes

import (
	"github.com/eastor112/clinic-grpc/controllers"
	"github.com/eastor112/clinic-grpc/middlewares"

	"github.com/gin-gonic/gin"
)

func PersonRouter(router *gin.Engine) {

	routes := router.Group("api/v1/persons")
	routes.POST("", controllers.PersonCreate)
	routes.GET("", controllers.PersonGet)
	routes.GET("/:id", controllers.PersonGetById)
	routes.PUT("/:id", controllers.PersonUpdate)
	routes.DELETE("/:id", controllers.PersonDelete)

}

func PatientRouter(router *gin.Engine) {

	routes := router.Group("api/v1/patients")
	routes.Use(middlewares.RouterProtection())
	routes.POST("", controllers.PatientCreate)
	routes.GET("", controllers.PatientGet)
	routes.GET("/:id", controllers.PatientGetById)
	routes.PATCH("/:id", controllers.PatientPatch)
	routes.DELETE("/:id", controllers.PatientDelete)

}
