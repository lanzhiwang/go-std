package main

/*
func HasPrefix(s, prefix string) bool
*/

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.HasPrefix("Gopher", "Go"))
	fmt.Println(strings.HasPrefix("Gopher", "C"))
	fmt.Println(strings.HasPrefix("Gopher", ""))
}

/*
$ go run HasPrefix.go
true
false
true
*/
