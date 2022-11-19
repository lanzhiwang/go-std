package main

/*
func TrimSpace(s string) string
*/

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.TrimSpace(" \t\n Hello, Gophers \n\t\r\n"))
}

/*
$ go run TrimSpace.go
Hello, Gophers
*/
