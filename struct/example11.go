package main

import "fmt"

type Drink struct {
	Name string
	Ice  bool
}

func main() {
	a := Drink{
		Name: "Lemonade",
		Ice:  true,
	}
	b := &a
	b.Ice = false
	if &a != b {
		fmt.Println("a and b are not the same")
	}
	fmt.Printf("%+v\n", *b)
	fmt.Printf("%+v\n", a)
	fmt.Printf("%p\n", b)
	fmt.Printf("%p\n", &a)
}

/*
{Name:Lemonade Ice:false}
{Name:Lemonade Ice:false}
0xc00000c030
0xc00000c030
*/
