package main

/*
func QuoteRuneToGraphic(r rune) string
*/

import (
	"fmt"
	"strconv"
)

func main() {
	s := strconv.QuoteRuneToGraphic('☺')
	fmt.Println(s)

	s = strconv.QuoteRuneToGraphic('\u263a')
	fmt.Println(s)

	s = strconv.QuoteRuneToGraphic('\u000a')
	fmt.Println(s)

	s = strconv.QuoteRuneToGraphic('	') // tab character
	fmt.Println(s)

}

/*
$ go run QuoteRuneToGraphic.go
'☺'
'☺'
'\n'
'\t'
*/

