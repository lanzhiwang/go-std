package main

// iota 只能被用在常量的赋值中，在每一个 const 关键字出现时，被重置为 0，然后每出现一个常量，iota 所代表的数值会自动增加1
// iota 可以理解成常量组中常量的计数器，不论该常量的值是什么，只要有一个常量，那么 iota 就加 1。

import "fmt"

const (
	a = iota
	b = iota
	c = iota
)

func main() {
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	fmt.Println("c: ", c)
}

/*
a:  0
b:  1
c:  2
*/
