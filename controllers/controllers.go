package controllers

import (
	"go-forum/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllThreads(c *gin.Context) {

	thread, err := models.GetThreadsDB()
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

	thread, err := models.GetThreadDB(threadID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, thread)

}

func CreateThread(c *gin.Context) {

	var newThread models.Threads
	err := c.ShouldBindJSON(&newThread)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	err = models.CreateThreadDB(newThread.Title, newThread.Description)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully Created the thread"})

}

func DeleteThreadByID(c *gin.Context) {

	thID := c.Param("thread_id")
	threadID, err := strconv.ParseInt(thID, 0, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = models.DeleteThreadDB(threadID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully Deleted the thread"})

}

func UpdateThread(c *gin.Context) {

	thID := c.Param("thread_id")
	threadID, err := strconv.ParseInt(thID, 0, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	thread, err := models.GetThreadDB(threadID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	err = c.ShouldBindJSON(&thread)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	err = models.UpdateThreadDB(thread)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, thread)

}

func GetAllPosts(c *gin.Context) {

	th := c.Param("thread_id")
	thID, err := strconv.ParseInt(th, 0, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	posts, err := models.GetPostsDB(thID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, posts)

}

func GetPostByID(c *gin.Context) {

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

	c.JSON(http.StatusOK, post)

}

func CreatePost(c *gin.Context) {

	var newPost models.Posts
	if err := c.ShouldBindJSON(&newPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	thID, err := strconv.ParseInt(c.Param("thread_id"), 0, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newPost.ThreadID = thID

	err = models.CreatePostDB(&newPost)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "post created successfully"})

}

func DeletePost(c *gin.Context) {

	th := c.Param("post_id")
	postID, err := strconv.ParseInt(th, 0, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.DeletePostDB(postID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successfully deleted"})

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
