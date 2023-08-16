package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	banner("Go", 6)
	banner("G☺", 6)

	s := "G☺!"
	for i := range s {
		fmt.Print(i, " ")
	}
	fmt.Println()
	for i, c := range s {
		//fmt.Print(i, ":", c, ",")
		fmt.Printf("%d: %c, ", i, c)
	}
}

/*
banner("Go", 6)

		Go
	  ------
*/
func banner(text string, width int) {
	// var padding = (len(text) - width) / 2
	// var padding
	// padding = (len(text) - width) / 2
	// padding := (width - len(text)) / 2
	padding := (width - utf8.RuneCountInString(text)) / 2
	for i := 0; i < padding; i++ {
		fmt.Print(" ")
	}
	fmt.Println(text)

	for i := 0; i < width; i++ {
		fmt.Print("-")
	}
	fmt.Println()
}
