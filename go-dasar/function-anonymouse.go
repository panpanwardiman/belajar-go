package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("Blacklist")
	} else {
		fmt.Println("Welcome")
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "anjing"
	}

	registerUser("panpan", blacklist)
	registerUser("anjing", func(name string) bool {
		return name == "anjing"
	})
}
