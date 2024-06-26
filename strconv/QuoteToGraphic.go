package main

/*
func QuoteToGraphic(s string) string
*/

import (
	"fmt"
	"strconv"
)

func main() {
	s := strconv.QuoteToGraphic("☺")
	fmt.Println(s)

	// This string literal contains a tab character.
	s = strconv.QuoteToGraphic("This is a \u263a	\u000a")
	fmt.Println(s)

	s = strconv.QuoteToGraphic(`" This is a ☺ \n "`)
	fmt.Println(s)

}

/*
$ go run QuoteToGraphic.go
"☺"
"This is a ☺\t\n"
"\" This is a ☺ \\n \""
*/
