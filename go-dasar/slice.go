package main

import "fmt"

func main() {
	month := [...]string{
		"januari",
		"februari",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}

	getMount := month[4:7]

	getMount = append(getMount, "2022")

	fmt.Println((getMount))
}
