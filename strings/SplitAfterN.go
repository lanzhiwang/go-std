package main

/*
func SplitAfterN(s, sep string, n int) []string
*/

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("%q\n", strings.SplitAfterN("a,b,c", ",", 2))
}

/*
$ go run SplitAfterN.go
["a," "b,c"]
*/
