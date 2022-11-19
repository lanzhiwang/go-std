package main

/*
func TrimSuffix(s, suffix string) string
*/

import (
	"fmt"
	"strings"
)

func main() {
	var s = "¡¡¡Hello, Gophers!!!"
	s = strings.TrimSuffix(s, ", Gophers!!!")
	s = strings.TrimSuffix(s, ", Marmots!!!")
	fmt.Print(s)
}

/*
$ go run TrimSuffix.go
¡¡¡Hello
*/
