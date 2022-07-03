package main

import "fmt"

func main() {
	name := "panpan"

	if name == "panpan" {
		fmt.Println(name)
	} else {
		fmt.Println("bukan panpan")
	}


	switch length := len(name) > length > 5 {
	case true:
		fmt.Println({"hello"})
	case false:
		fmt.println("world")
	default: 
		fmt.println("hello world")
	}


}