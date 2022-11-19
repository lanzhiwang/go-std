package main

/*
func Itoa(i int) string
*/

import (
	"fmt"
	"strconv"
)

func main() {
	i := 10
	s := strconv.Itoa(i)
	fmt.Printf("%T, %v\n", s, s)

}

/*
$ go run Itoa.go
string, 10
*/
