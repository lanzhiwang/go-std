package main

/*
func ReplaceAll(s, old, new string) string
*/

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.ReplaceAll("oink oink oink", "oink", "moo"))
}

/*
$ go run ReplaceAll.go
moo moo moo
*/
