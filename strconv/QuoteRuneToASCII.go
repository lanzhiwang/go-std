package main

/*
func QuoteRuneToASCII(r rune) string
*/

import (
	"fmt"
	"strconv"
)

func main() {
	s := strconv.QuoteRuneToASCII('☺')
	fmt.Println(s)

}

/*
$ go run QuoteRuneToASCII.go
'\u263a'
*/
