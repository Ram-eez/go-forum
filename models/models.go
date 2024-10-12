package models

import (
	"fmt"
	"go-forum/config"
	"time"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Threads struct {
	ID          int64   `json:"id" gorm:"autoIncrement"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Posts       []Posts `json:"posts" gorm:"foreignKey:ThreadID"`
}

type Posts struct {
	ID        int64  `json:"post_id" gorm:"autoIncrement;column:post_id"`
	Content   string `json:"content"`
	ThreadID  int64  `json:"thread_id"`
	CreatedAt string `json:"created_at"`
}

func init() {
	config.Connect()
	db = config.GetDB().Debug()
	if db == nil {
		fmt.Println("Database connection failed")
	} else {
		fmt.Println("Database connected successfully")
	}
	db.AutoMigrate(&Threads{}, &Posts{})
}

func CreateThread(title string, description string) error {

	newThread := Threads{
		Title:       title,
		Description: description,
	}
	result := db.Create(&newThread)

	if result.Error != nil {
		return result.Error
	}
	return nil

}

func GetAllThreads() ([]Threads, error) {

	var threads []Threads
	result := db.Preload("Posts").Find(&threads)
	if result.Error != nil {
		return nil, result.Error
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return threads, nil

}

func GetByID(id int64) (*Threads, error) {

	var getTread Threads
	result := db.Preload("Posts").Find(&getTread, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &getTread, nil

}

func DeleteThread(id int64) error {

	result := db.Where("id = ?", id).Delete(&Threads{})
	if result.Error != nil {
		return result.Error
	}

	return nil

}

func UpdateThread(thread *Threads) error {

	result := db.Model(thread).Updates(Threads{
		Title:       thread.Title,
		Description: thread.Description,
	})
	if result.Error != nil {
		return result.Error
	}

	return nil

}

func CreatePostDB(content string, threadID int64) (Posts, error) {

	newPost := Posts{
		Content:   content,
		ThreadID:  threadID,
		CreatedAt: time.Now().String(),
	}

	result := db.Create(&newPost)
	fmt.Println("New Post ID:", newPost.ID)
	if result.Error != nil {
		return Posts{}, result.Error
	}

	return newPost, nil

}

func DeletePostDB(id int64) error {

	result := db.Where("post_id = ?", id).Delete(&Posts{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
