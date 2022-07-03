package main

import (
	"errors"
	"fmt"
)

func pembagi(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Error Pembagi")
	} else {
		result := nilai / pembagi
		return result, nil
	}
}

func main() {
	var cthError = errors.New("oops error")
	fmt.Println(cthError)

	result, err := pembagi(10, 0)
	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err.Error())
	}
}
