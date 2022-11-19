package main

/*
func IsPrint(r rune) bool
*/

import (
	"fmt"
	"strconv"
)

func main() {
	c := strconv.IsPrint('\u263a')
	fmt.Println(c)

	bel := strconv.IsPrint('\007')
	fmt.Println(bel)

}

/*
$ go run IsPrint.go
true
false
*/
