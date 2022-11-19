package main

/*
func LastIndexByte(s string, c byte) int
*/

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.LastIndexByte("Hello, world", 'l'))
	fmt.Println(strings.LastIndexByte("Hello, world", 'o'))
	fmt.Println(strings.LastIndexByte("Hello, world", 'x'))
}

/*
$ go run LastIndexByte.go
10
8
-1
*/
