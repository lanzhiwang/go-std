package main

/*
func AppendQuoteRuneToASCII(dst []byte, r rune) []byte
*/

import (
	"fmt"
	"strconv"
)

func main() {
	b := []byte("rune (ascii):")
	b = strconv.AppendQuoteRuneToASCII(b, '☺')
	fmt.Println(string(b))

}

/*
$ go run AppendQuoteRuneToASCII.go
rune (ascii):'\u263a'
*/
