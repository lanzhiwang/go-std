package main

import (
	"fmt"
	"unicode"
)

func main() {
	const ucG = 'G'
	fmt.Printf("%#U\n", unicode.ToLower(ucG))

}

/*
$ go run ToLower.go
U+0067 'g'
*/
