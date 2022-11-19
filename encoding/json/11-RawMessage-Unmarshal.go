package main

/*
type RawMessage
func (m RawMessage) MarshalJSON() ([]byte, error)
func (m *RawMessage) UnmarshalJSON(data []byte) error
*/

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	type Color struct {
		Space string
		Point json.RawMessage // delay parsing until we know the color space
	}
	type RGB struct {
		R uint8
		G uint8
		B uint8
	}
	type YCbCr struct {
		Y  uint8
		Cb int8
		Cr int8
	}

	var j = []byte(`[
	{"Space": "YCbCr", "Point": {"Y": 255, "Cb": 0, "Cr": -10}},
	{"Space": "RGB",   "Point": {"R": 98, "G": 218, "B": 255}}
]`)
	var colors []Color
	err := json.Unmarshal(j, &colors)
	if err != nil {
		log.Fatalln("error:", err)
	}
	fmt.Println(colors)

	for _, c := range colors {
		var dst any
		switch c.Space {
		case "RGB":
			dst = new(RGB)
		case "YCbCr":
			dst = new(YCbCr)
		}
		err := json.Unmarshal(c.Point, dst)
		if err != nil {
			log.Fatalln("error:", err)
		}
		fmt.Println(c.Space, dst)
	}
}

/*
$ go run 10-RawMessage-Unmarshal.go
[
	{YCbCr [123 34 89 34 58 32 50 53 53 44 32 34 67 98 34 58 32 48 44 32 34 67 114 34 58 32 45 49 48 125]}
	{RGB [123 34 82 34 58 32 57 56 44 32 34 71 34 58 32 50 49 56 44 32 34 66 34 58 32 50 53 53 125]}
]
YCbCr &{255 0 -10}
RGB &{98 218 255}
*/
