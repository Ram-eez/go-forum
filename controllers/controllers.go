package controllers

import (
	"fmt"
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

	thID := c.Param("thread_id")
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

	thID := c.Param("thread_id")
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

	thID := c.Param("thread_id")
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

func CreatePost(c *gin.Context) {

	var newPost models.Posts
	if err := c.ShouldBindJSON(&newPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if newPost.ThreadID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "thread_id must be provided"})
		return
	}

	createdPost, err := models.CreatePostDB(newPost.Content, newPost.ThreadID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "post created successfully", "post_id": createdPost.ID})

}

func DeletePost(c *gin.Context) {
	th := c.Param("post_id")
	postID, err := strconv.ParseInt(th, 0, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("Attempting to delete PostID: %d\n", postID)

	if err := models.DeletePostDB(postID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successfully deleted"})
}

func GetPostByID(c *gin.Context) {
	th := c.Param("post_id")
	postID, err := strconv.ParseInt(th, 0, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	post, err := models.GetPostDB(postID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, post)
}

func UpadtePost(c *gin.Context) {
	th := c.Param("post_id")
	postID, err := strconv.ParseInt(th, 0, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post, err := models.GetPostDB(postID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.UpdatePostDB(post); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "post updates successfully"})
}
