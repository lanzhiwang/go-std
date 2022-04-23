package main

import "fmt"

type Point struct{
    x, y int
}

//乘2
func (p *Point) double() int {
	*p = Point{x: 2, y: 2}
    fmt.Printf("double p = %v\n", *p)
    return 0
}

//平方
func (p Point) square() int {
    p = Point{x: 3, y: 3}
    fmt.Printf("square p = %v\n", p)
    return 0
}

func main() {
    var i Point = Point{x: 1, y: 1}
    i.double()  // receiver 为对象的指针，原对象被修改
    fmt.Println("i = ", i)

    i.square()  //receiver 为对象的值，原对象不会被修改
    fmt.Println("i = ", i)
}

/*
$ go run 03.go
double p = {2 2}
i =  {2 2}
square p = {3 3}
i =  {2 2}
*/
