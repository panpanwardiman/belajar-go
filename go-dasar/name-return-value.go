package main

import "fmt"

func getNameFull() (first, last string) {
	first = "panpan"
	last = "wardiman"
	return
}

func main() {
	a, b := getNameFull()
	fmt.Println(a)
	fmt.Println(b)
}
