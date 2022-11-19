package main

/*
func Trim(s, cutset string) string
*/

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print(strings.Trim("¡¡¡Hello, Gophers!!!", "!¡"))
}

/*
$ go run Trim.go
Hello, Gophers
*/
