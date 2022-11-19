package main

/*
func QuoteRune(r rune) string
*/

import (
	"fmt"
	"strconv"
)

func main() {
	s := strconv.QuoteRune('☺')
	fmt.Println(s)

}

/*
$ go run QuoteRune.go
'☺'
*/
