package models

import (
	"fmt"
	"go-forum/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type User struct {
	ID       int64
	Username string
	Password string
}

type Threads struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func init() {
	config.Connect()
	db = config.GetDB().Debug()
	if db == nil {
		fmt.Println("Database connection failed")
	} else {
		fmt.Println("Database connected successfully")
	}
}

func CreateThread(title string, description string) error {
	newThread := Threads{
		Title:       title,
		Description: description,
	}
	result := db.Create(&newThread)

	// fmt.Println(result.Error)
	// fmt.Println(result.RowsAffected)
	if result.Error != nil {
		return result.Error
	}
	return nil

}

func GetAllThreads() ([]Threads, error) {
	var threads []Threads
	result := db.Find(&threads)
	if result.Error != nil {
		return nil, result.Error
	}
	return threads, nil
}

func GetByID(id int64) (*Threads, error) {
	var getTread Threads
	result := db.Find(&getTread, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &getTread, nil
}

func DeleteThread(id int64) error {
	var deleteThread Threads
	result := db.Delete(&deleteThread, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
