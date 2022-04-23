package main

import "fmt"

type Allergen int

const (
	IgEggs         Allergen = 1 << iota // 1 << 0 which is 00000001
	IgChocolate                         // 1 << 1 which is 00000010
	IgNuts                              // 1 << 2 which is 00000100
	IgStrawberries                      // 1 << 3 which is 00001000
	IgShellfish                         // 1 << 4 which is 00010000
)

func main() {
	fmt.Println("IgEggs: ", IgEggs)
	fmt.Println("IgChocolate: ", IgChocolate)
	fmt.Println("IgNuts: ", IgNuts)
	fmt.Println("IgStrawberries: ", IgStrawberries)
	fmt.Println("IgShellfish: ", IgShellfish)
}

/*
IgEggs:  1
IgChocolate:  2
IgNuts:  4
IgStrawberries:  8
IgShellfish:  16
*/
