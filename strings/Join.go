package main

/*
func Join(elems []string, sep string) string
*/

import (
	"fmt"
	"strings"
)

func main() {
	s := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(s, ", "))
}

/*
$ go run Join.go
foo, bar, baz
*/
