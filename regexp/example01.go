package main

import (
	"fmt"
	"log"
	"regexp"
)

func main() {
	needle := "chocolate"
	haystack := "Chocolate is my favorite!"

	// func MatchString(pattern string, s string) (matched bool, err error)
	match, err := regexp.MatchString(needle, haystack)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(match)  // false
}
