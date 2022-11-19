package main

/*
func TrimPrefix(s, prefix string) string
*/

import (
	"fmt"
	"strings"
)

func main() {
	var s = "¡¡¡Hello, Gophers!!!"
	s = strings.TrimPrefix(s, "¡¡¡Hello, ")
	s = strings.TrimPrefix(s, "¡¡¡Howdy, ")
	fmt.Print(s)
}

/*
$ go run TrimPrefix.go
Gophers!!!
*/
