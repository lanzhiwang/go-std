package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Switch struct {
	On bool `json:"on"`
}

func main() {
	jsonStringData := `{"on":true}`
	jsonByteData := []byte(jsonStringData)
	s := Switch{}
	err := json.Unmarshal(jsonByteData, &s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", s)

	jsonStringData = `{"on":"true"}`
	jsonByteData = []byte(jsonStringData)
	s = Switch{}
	err = json.Unmarshal(jsonByteData, &s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", s)
}

/*
{On:true}
2022/04/23 14:34:43 json: cannot unmarshal string into Go struct field Switch.on of type bool
exit status 1
*/
