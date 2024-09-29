package models

import (
	"fmt"
	"go-forum/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type User struct {
	ID       int
	Username string
	Password string
}

type Threads struct {
	ID          int
	Title       string
	Description string
}

// // func init() {
// // 	config.Connect()
// // 	db = config.GetDB()
// }

func CreateThread(title string, description string) {
	config.Connect()
	db = config.GetDB()
	newThread := Threads{
		Title:       title,
		Description: description,
	}
	result := db.Create(&newThread)

	fmt.Println(result.Error)
	fmt.Println(result.RowsAffected)
	fmt.Println(newThread.ID)
}
