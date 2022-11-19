package main

/*
func AppendFloat(dst []byte, f float64, fmt byte, prec, bitSize int) []byte
*/

import (
	"fmt"
	"strconv"
)

func main() {
	b32 := []byte("float32:")
	fmt.Println(b32)
	b32 = strconv.AppendFloat(b32, 3.1415926535, 'E', -1, 32)
	fmt.Println(string(b32))
	fmt.Println(b32)

	b64 := []byte("float64:")
	fmt.Println(b64)
	b64 = strconv.AppendFloat(b64, 3.1415926535, 'E', -1, 64)
	fmt.Println(string(b64))
	fmt.Println(b64)

}

/*
$ go run AppendFloat.go
[102 108 111 97 116 51 50 58]
float32:3.1415927E+00
[102 108 111 97 116 51 50 58 51 46 49 52 49 53 57 50 55 69 43 48 48]
[102 108 111 97 116 54 52 58]
float64:3.1415926535E+00
[102 108 111 97 116 54 52 58 51 46 49 52 49 53 57 50 54 53 51 53 69 43 48 48]
*/
