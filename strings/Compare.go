package main

/*
func Compare(a, b string) int
*/

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Compare("a", "b"))
	fmt.Println(strings.Compare("a", "a"))
	fmt.Println(strings.Compare("b", "a"))
}

/*
$ go run Compare.go
-1
0
1
*/
