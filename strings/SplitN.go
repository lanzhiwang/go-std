package main

/*
func SplitN(s, sep string, n int) []string
*/

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("%q\n", strings.SplitN("a,b,c", ",", 2))
	z := strings.SplitN("a,b,c", ",", 0)
	fmt.Printf("%q (nil = %v)\n", z, z == nil)
}

/*
$ go run SplitN.go
["a" "b,c"]
[] (nil = true)
*/
