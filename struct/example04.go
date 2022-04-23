package main

import (
	"fmt"
	"reflect"
)

type Movie struct {
	Name   string
	Rating float32
}

func main() {
	m := new(Movie)
	fmt.Printf("%+v\n", m)
	fmt.Printf("%T\n", m)
	fmt.Println(reflect.TypeOf(m).String())
	m.Name = "Metropolis"
	m.Rating = 0.99
	fmt.Printf("%+v\n", m)
}

/*
&{Name: Rating:0}
*main.Movie
*main.Movie
&{Name:Metropolis Rating:0.99}
*/
