package main

/*
func Replace(s, old, new string, n int) string
*/

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))
}

/*
$ go run Replace.go
oinky oinky oink
moo moo moo
*/
