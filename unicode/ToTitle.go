package main

import (
	"fmt"
	"unicode"
)

func main() {
	const ucG = 'g'
	fmt.Printf("%#U\n", unicode.ToTitle(ucG))

}

/*
$ go run ToTitle.go
U+0047 'G'
*/
