package main

import "fmt"

type Point struct{
    x, y int
}

//乘2
func double(x *Point) int {
	*x = Point{x: 2, y: 2}
    fmt.Printf("double p = %s\n", *x)
    return 0
}

//平方
func square(x Point) int {
    x = Point{x: 3, y: 3}
    fmt.Printf("square p = %s\n", x)
    return 0
}

func main() {
    var i Point = Point{x: 1, y: 1}
    double(&i)  // receiver 为对象的指针，原对象被修改
    fmt.Println("i = ", i)

    square(i)  //receiver 为对象的值，原对象不会被修改
    fmt.Println("i = ", i)
}
