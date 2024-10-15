package routes

import (
	"go-forum/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterThreadRoutes(router *gin.Engine) {
	ThreadRoutes := router.Group("/threads")
	{
		ThreadRoutes.GET("/", controllers.GetAllThreads)
		ThreadRoutes.GET("/:thread_id", controllers.GetThreadByID)
		ThreadRoutes.DELETE("/:thread_id", controllers.DeleteThreadByID)
		ThreadRoutes.POST("/", controllers.CreateThread)
		ThreadRoutes.PUT("/:thread_id", controllers.UpdateThread)

		ThreadRoutes.POST("/:thread_id/posts/", controllers.CreatePost)
		ThreadRoutes.DELETE("/:thread_id/posts/:post_id", controllers.DeletePost)
		ThreadRoutes.GET("/:thread_id/posts/:post_id", controllers.GetPostByID)
		ThreadRoutes.PUT("/:thread_id/posts/:post_id", controllers.UpadtePost)
		ThreadRoutes.GET("/:thread_id/posts/", controllers.GetAllPosts)

	}

	UserRoutes := router.Group("/users")
	{
		UserRoutes.POST("/", controllers.CreateUser)
		UserRoutes.GET("/", controllers.GetAllUsers)
		UserRoutes.GET("/:user_id/", controllers.GetUserByID)
		UserRoutes.DELETE("/:user_id/", controllers.DeleteUser)
	}

}
