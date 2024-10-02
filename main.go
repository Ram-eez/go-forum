package main

import (
	"fmt"
	"go-forum/models"
)

func main() {

	//models.CreateThread("abc", "5")

	//allthreads := models.GetAllThreads()

	deleted := models.DeleteThread(3)

	fmt.Println(deleted)

}
