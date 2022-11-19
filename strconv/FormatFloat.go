package main

/*
func FormatFloat(f float64, fmt byte, prec, bitSize int) string
*/

import (
	"fmt"
	"strconv"
)

func main() {
	v := 3.1415926535

	s32 := strconv.FormatFloat(v, 'E', -1, 32)
	fmt.Printf("%T, %v\n", s32, s32)

	s64 := strconv.FormatFloat(v, 'E', -1, 64)
	fmt.Printf("%T, %v\n", s64, s64)

}

/*
$ go run FormatFloat.go
string, 3.1415927E+00
string, 3.1415926535E+00
*/
