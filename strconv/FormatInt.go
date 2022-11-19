package main

/*
func FormatInt(i int64, base int) string
*/

import (
	"fmt"
	"strconv"
)

func main() {
	v := int64(-42)

	s10 := strconv.FormatInt(v, 10)
	fmt.Printf("%T, %v\n", s10, s10)

	s16 := strconv.FormatInt(v, 16)
	fmt.Printf("%T, %v\n", s16, s16)

}

/*
$ go run FormatInt.go
string, -42
string, -2a
*/
