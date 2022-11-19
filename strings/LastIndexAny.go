package main

/*
func LastIndexAny(s, chars string) int
*/

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.LastIndexAny("go gopher", "go"))
	fmt.Println(strings.LastIndexAny("go gopher", "rodent"))
	fmt.Println(strings.LastIndexAny("go gopher", "fail"))
}

/*
$ go run LastIndexAny.go
4
8
-1
*/
