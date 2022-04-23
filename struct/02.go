package main

import "fmt"

type Sli []string

//乘2
func (p *Sli) double() int {
	*p = append(*p, "double1")
    fmt.Printf("double p = %s\n", *p)
    return 0
}

//平方
func (p Sli) square() int {
    p = append(p, "double2")
    fmt.Printf("square p = %s\n", p)
    return 0
}

func main() {
    var i Sli = []string{"double3"}
    i.double()  // receiver 为对象的指针，原对象被修改
    fmt.Println("i = ", i)

    i.square()  //receiver 为对象的值，原对象不会被修改
    fmt.Println("i = ", i)
}

/*
$ go run 02.go
double p = [double3 double1]
i =  [double3 double1]
square p = [double3 double1 double2]
i =  [double3 double1]
*/
