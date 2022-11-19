package main

/*
func Repeat(s string, count int) string
*/

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("ba" + strings.Repeat("na", 2))
}

/*
$ go run Repeat.go
banana
*/
