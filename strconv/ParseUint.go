package main

/*
func ParseUint(s string, base int, bitSize int) (uint64, error)
*/

import (
	"fmt"
	"strconv"
)

func main() {
	v := "42"
	if s, err := strconv.ParseUint(v, 10, 32); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}
	if s, err := strconv.ParseUint(v, 10, 64); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}

}

/*
$ go run ParseUint.go
uint64, 42
uint64, 42
*/
