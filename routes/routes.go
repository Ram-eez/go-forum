package routes

import (
	"go-forum/controllers"
	"go-forum/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterThreadRoutes(router *gin.Engine) {

	UserRoutes := router.Group("/users")
	{
		UserRoutes.POST("/signup", controllers.Signup) // Public route
		UserRoutes.POST("/login", controllers.Login)   // Public route
	}

	protected := router.Group("/")
	protected.Use(middleware.VerifyToken) // Protect these routes
	{
		// Protected thread routes
		ThreadRoutes := protected.Group("/threads")
		{
			ThreadRoutes.GET("/", controllers.GetAllThreads)
			ThreadRoutes.GET("/:thread_id", controllers.GetThreadByID)
			ThreadRoutes.DELETE("/:thread_id", controllers.DeleteThreadByID)
			ThreadRoutes.POST("/", controllers.CreateThread)
			ThreadRoutes.PUT("/:thread_id", controllers.UpdateThread)

			ThreadRoutes.POST("/:thread_id/posts", controllers.CreatePost)
			ThreadRoutes.DELETE("/:thread_id/posts/:post_id", controllers.DeletePost)
			ThreadRoutes.GET("/:thread_id/posts/:post_id", controllers.GetPostByID)
			ThreadRoutes.PUT("/:thread_id/posts/:post_id", controllers.UpadtePost)
			ThreadRoutes.GET("/:thread_id/posts", controllers.GetAllPosts)
		}
		UserRoutes.GET("/", controllers.GetAllUsers)           // Protected
		UserRoutes.GET("/:user_id", controllers.GetUserByID)   // Protected
		UserRoutes.DELETE("/:user_id", controllers.DeleteUser) // Protected
		UserRoutes.PUT("/:user_id", controllers.UpdateUser)    // Protected
		UserRoutes.GET("/validate", controllers.Validate)
	}
}
