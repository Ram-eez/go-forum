package routes

import (
	"go-forum/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterThreadRoutes(router *gin.Engine) {
	ThreadRoutes := router.Group("/threads")
	{
		ThreadRoutes.GET("/", controllers.GetThreads)
		ThreadRoutes.GET("/:thread_id", controllers.GetThreadByID)
		ThreadRoutes.DELETE("/:thread_id", controllers.DeleteThreadByID)
		ThreadRoutes.POST("/", controllers.CreateAThread)
		ThreadRoutes.PUT("/:thread_id", controllers.UpdateAThread)

		ThreadRoutes.POST("/:thread_id/posts", controllers.CreatePost)
		ThreadRoutes.DELETE("/:thread_id/posts/:post_id", controllers.DeletePost)
		ThreadRoutes.GET("/:thread_id/posts/:post_id", controllers.GetPostByID)
	}

}
