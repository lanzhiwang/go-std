package main

import "fmt"

const (
	i = iota
	j = 3.14
	k = iota
	l
)

func main() {
	fmt.Println("i: ", i)
	fmt.Println("j: ", j)
	fmt.Println("k: ", k)
	fmt.Println("l: ", l)
}

/*
i:  0
j:  3.14
k:  2
l:  3
*/
