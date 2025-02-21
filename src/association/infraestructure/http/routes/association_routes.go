package routes

import (
	"github.com/gin-gonic/gin"

	"api/src/association/infraestructure/http/controllers"
)

func SetupRoutes(router *gin.Engine, controller *controllers.AssociationController) {
	association := router.Group("/associations")
	{
		association.POST("/", controller.Create)
		association.GET("/", controller.GetAll)
		association.GET("/:id", controller.GetByID)
		association.PUT("/:id", controller.Update)
		association.DELETE("/:id", controller.Delete)
	}
}
