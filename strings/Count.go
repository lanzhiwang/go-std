package main

/*
func Count(s, substr string) int
*/

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Count("cheese", "e"))
	fmt.Println(strings.Count("five", "")) // before & after each rune
}

/*
$ go run Count.go
3
5
*/