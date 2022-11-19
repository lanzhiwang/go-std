package main

/*
func IndexRune(s string, r rune) int
*/

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.IndexRune("chicken", 'k'))
	fmt.Println(strings.IndexRune("chicken", 'd'))
}

/*
$ go run IndexRune.go
4
-1
*/
