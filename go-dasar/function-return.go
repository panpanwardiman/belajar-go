package main

import "fmt"

func main() {
	result := getName("panpan")
	fmt.Println(result)
}

func getName(name string) string {
	return "Hello " + name
}