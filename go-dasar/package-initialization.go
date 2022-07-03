package main

import (
	"fmt"
	"belajar-go/go-dasar/database"
)

func main() {
	result := database.GetConnection()
	fmt.Println(result)
}