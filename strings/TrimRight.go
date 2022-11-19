package main

/*
func TrimRight(s, cutset string) string
*/

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print(strings.TrimRight("¡¡¡Hello, Gophers!!!", "!¡"))
}

/*
$ go run TrimRight.go
¡¡¡Hello, Gophers
*/
