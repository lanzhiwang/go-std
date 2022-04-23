package main

import "fmt"

const (
	Apple, Banana = iota + 1, iota + 2
	Cherimoya, Durian
	Elderberry, Fig
)

func main() {
	fmt.Println("Apple: ", Apple)
	fmt.Println("Banana: ", Banana)
	fmt.Println("Cherimoya: ", Cherimoya)
	fmt.Println("Durian: ", Durian)
	fmt.Println("Elderberry: ", Elderberry)
	fmt.Println("Fig: ", Fig)
}

/*
Apple:  1
Banana:  2
Cherimoya:  2
Durian:  3
Elderberry:  3
Fig:  4
*/
