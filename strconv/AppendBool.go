package main

/*
func AppendBool(dst []byte, b bool) []byte
*/

import (
	"fmt"
	"strconv"
)

func main() {
	b := []byte("bool:")
	fmt.Println(b)
	b = strconv.AppendBool(b, true)
	fmt.Println(string(b))
	fmt.Println(b)

}

/*
$ go run AppendBool.go
[98 111 111 108 58]
bool:true
[98 111 111 108 58 116 114 117 101]
*/
