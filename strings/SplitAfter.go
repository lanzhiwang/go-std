package main

/*
func SplitAfter(s, sep string) []string
*/

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("%q\n", strings.SplitAfter("a,b,c", ","))
}

/*
$ go run SplitAfter.go
["a," "b," "c"]
*/
