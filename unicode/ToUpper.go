package main

import (
	"fmt"
	"unicode"
)

func main() {
	const ucG = 'g'
	fmt.Printf("%#U\n", unicode.ToUpper(ucG))

}

/*
$ go run ToUpper.go
U+0047 'G'
*/
