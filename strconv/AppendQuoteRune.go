package main

/*
func AppendQuoteRune(dst []byte, r rune) []byte
*/

import (
	"fmt"
	"strconv"
)

func main() {
	b := []byte("rune:")
	b = strconv.AppendQuoteRune(b, '☺')
	fmt.Println(string(b))

}

/*
$ go run AppendQuoteRune.go
rune:'☺'
*/
