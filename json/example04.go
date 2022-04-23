package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name    string   `json:"name,omitempty"`
	Age     int      `json:"age,omitempty"`
	Hobbies []string `json:"hobbies,omitempty"`
}

func main() {
	hobbies := []string{"Cycling", "Cheese", "Techno"}
	p := Person{
		Name:    "George",
		Age:     40,
		Hobbies: hobbies,
	}
	fmt.Printf("%+v\n", p)
	jsonByteData, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}
	jsonStringData := string(jsonByteData)
	fmt.Println(jsonStringData)

	jsonByteData, err = json.Marshal(Person{})
	if err != nil {
		log.Fatal(err)
	}
	jsonStringData = string(jsonByteData)
	fmt.Println(jsonStringData)
}

/*
{Name:George Age:40 Hobbies:[Cycling Cheese Techno]}
{"name":"George","age":40,"hobbies":["Cycling","Cheese","Techno"]}
{}
*/
