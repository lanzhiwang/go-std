package main

/*
func IndexAny(s, chars string) int
*/

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.IndexAny("chicken", "aeiouy"))
	fmt.Println(strings.IndexAny("crwth", "aeiouy"))
}

/*
$ go run IndexAny.go
2
-1
*/
