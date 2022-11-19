package main

/*
func AppendQuote(dst []byte, s string) []byte
*/

import (
	"fmt"
	"strconv"
)

func main() {
	b := []byte("quote:")
	b = strconv.AppendQuote(b, `"Fran & Freddie's Diner"`)
	fmt.Println(string(b))

}

/*
$ go run AppendQuote.go
quote:"\"Fran & Freddie's Diner\""
*/
