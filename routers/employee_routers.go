package routers

import (
	"awesomeProject/controllers"
	"github.com/gin-gonic/gin"
)

func SetupEmployeeRoutes(r *gin.RouterGroup) {
	category := r.Group("/employee")
	{
		category.POST("/", controllers.CreateEmployee)
	}
}
