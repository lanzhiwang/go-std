package main

import "fmt"

type Integer int

//乘2
func (p *Integer) double() int {
    *p = *p * 2
    fmt.Printf("double p = %d\n", *p)
    return 0
}

//平方
func (p Integer) square() int {
    p = p * p
    fmt.Printf("square p = %d\n", p)
    return 0
}

func main() {
    var i Integer = 2
    i.double()  // receiver 为对象的指针，原对象被修改
    fmt.Println("i = ", i)

    i.square()  //receiver 为对象的值，原对象不会被修改
    fmt.Println("i = ", i)
}

/*
$ go run 01.go
double p = 4
i =  4
square p = 16
i =  4
*/
