package routes

import (
	"go-forum/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterThreadRoutes(router *gin.Engine) {
	ThreadRoutes := router.Group("/threads")
	{
		ThreadRoutes.GET("/", controllers.GetThreads)
		ThreadRoutes.GET("/:ID", controllers.GetThreadByID)
		ThreadRoutes.DELETE("/:ID", controllers.DeleteThreadByID)
		ThreadRoutes.POST("/", controllers.CreateAThread)
		ThreadRoutes.PUT("/:ID", controllers.UpdateAThread)

		ThreadRoutes.POST("/:threadID/posts	", controllers.CreatePost)
	}

}
