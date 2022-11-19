package main

/*
func AppendInt(dst []byte, i int64, base int) []byte
*/

import (
	"fmt"
	"strconv"
)

func main() {
	b10 := []byte("int (base 10):")
	fmt.Println(b10)
	b10 = strconv.AppendInt(b10, -42, 10)
	fmt.Println(string(b10))
	fmt.Println(b10)

	b16 := []byte("int (base 16):")
	fmt.Println(b16)
	b16 = strconv.AppendInt(b16, -42, 16)
	fmt.Println(string(b16))
	fmt.Println(b16)

}

/*
$ go run AppendInt.go
[105 110 116 32 40 98 97 115 101 32 49 48 41 58]
int (base 10):-42
[105 110 116 32 40 98 97 115 101 32 49 48 41 58 45 52 50]
[105 110 116 32 40 98 97 115 101 32 49 54 41 58]
int (base 16):-2a
[105 110 116 32 40 98 97 115 101 32 49 54 41 58 45 50 97]
*/
