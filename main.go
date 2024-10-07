package main

import (
	"go-forum/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	routes.RegisterThreadRoutes(router)

	router.Run(":8080")
}
