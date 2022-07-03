package main

import "fmt"

func sayHelloFilter(name string, filter func(string) string) {
	nameFilter := filter(name)
	fmt.Println(nameFilter)
}

func doFilter(name string) string {
	if name == "anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloFilter("panpan", doFilter)
}
