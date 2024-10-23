package models

import (
	"fmt"
	"go-forum/config"
	"time"

	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	ID       int64     `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Threads  []Threads `json:"threads" gorm:"foreignKey:UserID"`
	Posts    []Posts   `json:"posts" gorm:"foreignKey:UserID"`
}

type Threads struct {
	ID          int64   `json:"id" gorm:"autoIncrement"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	UserID      int64   `json:"user_id"`
	Posts       []Posts `json:"posts" gorm:"foreignKey:ThreadID"`
}

type Posts struct {
	ID        int64  `json:"post_id" gorm:"autoIncrement;column:post_id"`
	Content   string `json:"content"`
	ThreadID  int64  `json:"thread_id"`
	UserID    int64  `json:"user_id"`
	CreatedAt string `json:"created_at"`
}

func init() {
	// config.LoadEnvVariables()
	config.Connect()
	db = config.GetDB()
	if db == nil {
		fmt.Println("Database connection failed")
	} else {
		fmt.Println("Database connected successfully")
	}
	db.AutoMigrate(&Threads{}, &Posts{})
}

func CreateUserDB(newUser User) error {

	if err := db.Create(&newUser).Error; err != nil {
		return err
	}
	return nil
}

func GetUsersDB() (*[]User, error) {
	var users []User
	if err := db.Preload("Threads.Posts").Preload("Posts").Omit("password").Find(&users).Error; err != nil {
		return nil, err
	}
	return &users, nil
}

func GetUserDB(id int64) (*User, error) {
	var user User
	if err := db.Preload("Threads.Posts").Preload("Posts").Omit("password").First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByEmail(email string) (*User, error) {
	var user User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func DeleteUserDB(id int64) error {
	var user User
	if err := db.Delete(&user, id).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUserDB(user *User) error {
	var existingUser User
	if err := db.First(&existingUser, user.ID).Error; err != nil {
		return err
	}

	if err := db.Model(user).Updates(User{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}).Error; err != nil {
		return err
	}
	return nil
}

func CreateThreadDB(newThread Threads) error {

	if err := db.Create(&newThread).Error; err != nil {
		return err
	}
	return nil

}

func GetThreadsDB() ([]Threads, error) {

	var threads []Threads
	result := db.Preload("Posts").Find(&threads)
	if result.Error != nil {
		return nil, result.Error
	}
	return threads, nil

}

func GetThreadDB(id int64) (*Threads, error) {

	var getTread Threads
	result := db.Preload("Posts").Find(&getTread, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &getTread, nil

}

func DeleteThreadDB(id int64) error {

	result := db.Where("id = ?", id).Delete(&Threads{})
	if result.Error != nil {
		return result.Error
	}

	return nil

}

func UpdateThreadDB(thread *Threads) error {

	var existingThread Threads
	if err := db.First(&existingThread, thread.ID).Error; err != nil {
		return err
	}

	result := db.Model(&existingThread).Updates(Threads{
		Title:       thread.Title,
		Description: thread.Description,
	})
	if result.Error != nil {
		return result.Error
	}

	return nil

}

func CreatePostDB(post *Posts) error {

	newPost := Posts{
		Content:   post.Content,
		ThreadID:  post.ThreadID,
		UserID:    post.UserID,
		CreatedAt: time.Now().String(),
	}

	result := db.Create(&newPost)
	fmt.Println("New Post ID:", newPost.ID)
	if result.Error != nil {
		return result.Error
	}

	return nil

}

func GetPostsDB(id int64) ([]Posts, error) {
	var posts []Posts

	if err := db.Where("thread_id = ?", id).Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func GetPostDB(id int64) (*Posts, error) {
	var post Posts
	result := db.Where("post_id = ?", id).Find(&post).Error
	if result != nil {
		return nil, result
	}

	return &post, nil
}

func DeletePostDB(id int64) error {

	result := db.Where("post_id = ?", id).Delete(&Posts{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func UpdatePostDB(post *Posts) error {

	if err := db.Model(&post).Updates(Posts{
		Content: post.Content,
	}); err != nil {
		return err.Error
	}
	return nil
}
