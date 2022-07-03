package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	fmt.Println(sumAll(10, 20, 30, 40))

	numbers := []int{20, 20, 20}
	fmt.Println(sumAll(numbers...))
}
