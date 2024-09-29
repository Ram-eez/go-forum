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

func CreateThread(title string, description string) *Threads {
	newThread := Threads{
		Title:       title,
		Description: description,
	}
	result := db.Create(&newThread)

	fmt.Println(result.Error)
	fmt.Println(result.RowsAffected)
	return &newThread

}

func GetAllThreads() []Threads {
	var threads []Threads
	db.Find(&threads)
	return threads
}

func GetByID(id int64) *Threads {
	var getTread Threads
	db.Find(&getTread, id)
	return &getTread
}

// func DeleteThread(id int64) *Threads {

// }
