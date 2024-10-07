package controllers

import (
	"go-forum/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetThreads(c *gin.Context) {
	thread, err := models.GetAllThreads()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, thread)
}
