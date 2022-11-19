package main

/*
func AppendQuoteToASCII(dst []byte, s string) []byte
*/

import (
	"fmt"
	"strconv"
)

func main() {
	b := []byte("quote (ascii):")
	b = strconv.AppendQuoteToASCII(b, `"Fran & Freddie's Diner"`)
	fmt.Println(string(b))

}

/*
$ go run AppendQuoteToASCII.go
quote (ascii):"\"Fran & Freddie's Diner\""
*/
