package main

/*
func AppendUint(dst []byte, i uint64, base int) []byte
*/

import (
	"fmt"
	"strconv"
)

func main() {
	b10 := []byte("uint (base 10):")
	b10 = strconv.AppendUint(b10, 42, 10)
	fmt.Println(string(b10))

	b16 := []byte("uint (base 16):")
	b16 = strconv.AppendUint(b16, 42, 16)
	fmt.Println(string(b16))

}

/*
$ go run AppendUint.go
uint (base 10):42
uint (base 16):2a
*/
