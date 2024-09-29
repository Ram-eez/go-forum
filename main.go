package main

import (
	"fmt"
	"go-forum/models"
)

func main() {

	models.CreateThread("abc", "third")

	allthreads := models.GetAllThreads()

	fmt.Println(allthreads)

}
