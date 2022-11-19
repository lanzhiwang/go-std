package main

/*
func ContainsRune(s string, r rune) bool
*/

import (
	"fmt"
	"strings"
)

func main() {
	// Finds whether a string contains a particular Unicode code point.
	// The code point for the lowercase letter "a", for example, is 97.
	fmt.Println(strings.ContainsRune("aardvark", 97))
	fmt.Println(strings.ContainsRune("timeout", 97))
}

/*
$ go run ContainsRune.go
true
false
*/
