package main

/*
func TrimLeft(s, cutset string) string
*/

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print(strings.TrimLeft("¡¡¡Hello, Gophers!!!", "!¡"))
}

/*
$ go run TrimLeft.go
Hello, Gophers!!!
*/
