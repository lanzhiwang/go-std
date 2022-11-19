package main

/*
func Quote(s string) string
*/

import (
	"fmt"
	"strconv"
)

func main() {
	// This string literal contains a tab character.
	s := strconv.Quote(`"Fran & Freddie's Diner	☺"`)
	fmt.Println(s)

}

/*
$ go run Quote.go
"\"Fran & Freddie's Diner\t☺\""
*/
