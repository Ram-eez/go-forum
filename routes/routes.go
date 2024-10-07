package routes

import (
	"go-forum/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterThreadRoutes(router *gin.Engine) {
	ThreadRoutes := router.Group("/threads")
	{
		ThreadRoutes.GET("/", controllers.GetThreads)
	}
}
