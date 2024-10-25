package routes

import (
	"go-forum/controllers"
	"go-forum/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterThreadRoutes(router *gin.Engine) {

	//Public routes
	router.POST("/signup", controllers.Signup)
	router.POST("/login", controllers.Login)
	router.GET("/threads", controllers.GetAllThreads)
	router.GET("/threads/:thread_id", controllers.GetThreadByID)

	authorized := router.Group("/")
	authorized.Use(middleware.VerifyToken)
	{
		//private routes
		threads := authorized.Group("/threads")
		{
			threads.POST("/", controllers.CreateThread)
			threads.PUT("/:thread_id", controllers.UpdateThread)
			threads.DELETE("/:thread_id", controllers.DeleteThreadByID)

			posts := threads.Group("/:thread_id/posts")
			{
				posts.POST("/", controllers.CreatePost)
				posts.GET("/", controllers.GetAllPosts)
				posts.GET("/:post_id", controllers.GetPostByID)
				posts.PUT("/:post_id", controllers.UpdatePost)
				posts.DELETE("/:post_id", controllers.DeletePost)
			}
		}

		users := authorized.Group("/users")
		{
			users.GET("/", controllers.GetAllUsers)
			users.GET("/validate", controllers.Validate)
			users.GET("/user", controllers.GetUserByID)
			users.PUT("/update", controllers.UpdateUser)
			users.DELETE("/delete", controllers.DeleteUser)
		}
	}

}
