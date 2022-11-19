package main

/*
func FormatUint(i uint64, base int) string
*/

import (
	"fmt"
	"strconv"
)

func main() {
	v := uint64(42)

	s10 := strconv.FormatUint(v, 10)
	fmt.Printf("%T, %v\n", s10, s10)

	s16 := strconv.FormatUint(v, 16)
	fmt.Printf("%T, %v\n", s16, s16)

}

/*
$ go run FormatUint.go
string, 42
string, 2a
*/
