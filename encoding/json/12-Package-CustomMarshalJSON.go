package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type Animal int

const (
	Unknown Animal = iota
	Gopher
	Zebra
)

func (a *Animal) UnmarshalJSON(b []byte) error {
	fmt.Println("UnmarshalJSON")
	fmt.Println(string(b))
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch strings.ToLower(s) {
	default:
		*a = Unknown
	case "gopher":
		*a = Gopher
	case "zebra":
		*a = Zebra
	}

	return nil
}

func (a Animal) MarshalJSON() ([]byte, error) {
	fmt.Println("MarshalJSON")
	var s string
	switch a {
	default:
		s = "unknown"
	case Gopher:
		s = "gopher"
	case Zebra:
		s = "zebra"
	}

	return json.Marshal(s)
}

func main() {
	blob := `["gopher","armadillo","zebra","unknown","gopher","bee","gopher","zebra"]`
	var zoo []Animal
	fmt.Println(zoo)
	if err := json.Unmarshal([]byte(blob), &zoo); err != nil {
		log.Fatal(err)
	}
	fmt.Println(zoo)

	census := make(map[Animal]int)
	for _, animal := range zoo {
		census[animal] += 1
	}

	fmt.Printf("Zoo Census:\n* Gophers: %d\n* Zebras:  %d\n* Unknown: %d\n",
		census[Gopher], census[Zebra], census[Unknown])

}

/*
$ go run 12-Package-CustomMarshalJSON.go
[]
UnmarshalJSON
"gopher"
UnmarshalJSON
"armadillo"
UnmarshalJSON
"zebra"
UnmarshalJSON
"unknown"
UnmarshalJSON
"gopher"
UnmarshalJSON
"bee"
UnmarshalJSON
"gopher"
UnmarshalJSON
"zebra"
[1 0 2 0 1 0 1 2]
Zoo Census:
* Gophers: 3
* Zebras:  2
* Unknown: 3
*/
