package main

/*
func IsGraphic(r rune) bool
*/

import (
	"fmt"
	"strconv"
)

func main() {
	shamrock := strconv.IsGraphic('☘')
	fmt.Println(shamrock)

	a := strconv.IsGraphic('a')
	fmt.Println(a)

	bel := strconv.IsGraphic('\007')
	fmt.Println(bel)

}

/*
$ go run IsGraphic.go
true
true
false
*/
