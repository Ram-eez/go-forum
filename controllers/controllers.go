package controllers

import (
	"go-forum/models"
	"net/http"
	"strconv"

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

func GetThreadByID(c *gin.Context) {

	thID := c.Param("ID")
	threadID, err := strconv.ParseInt(thID, 0, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	thread, err := models.GetByID(threadID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, thread)

}

func DeleteThreadByID(c *gin.Context) {

	thID := c.Param("ID")
	threadID, err := strconv.ParseInt(thID, 0, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = models.DeleteThread(threadID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully Deleted the thread"})

}

func CreateAThread(c *gin.Context) {

	var newThread models.Threads
	err := c.ShouldBindJSON(&newThread)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	err = models.CreateThread(newThread.Title, newThread.Description)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully Created the thread"})

}

func UpdateAThread(c *gin.Context) {

	thID := c.Param("ID")
	threadID, err := strconv.ParseInt(thID, 0, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	thread, err := models.GetByID(threadID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	err = c.ShouldBindJSON(&thread)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	err = models.UpdateThread(thread)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, thread)

}
