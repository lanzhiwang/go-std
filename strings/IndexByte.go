package main

/*
func IndexByte(s string, c byte) int
*/

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.IndexByte("golang", 'g'))
	fmt.Println(strings.IndexByte("gophers", 'h'))
	fmt.Println(strings.IndexByte("golang", 'x'))
}

/*
$ go run IndexByte.go
0
3
-1
*/
