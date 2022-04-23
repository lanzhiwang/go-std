package main

import "fmt"

type AudioOutput int

const (
	OutMute AudioOutput = iota
	OutMono
	OutStereo
	_
	_
	OutSurround
)

func main() {
	fmt.Println("OutMute: ", OutMute)
	fmt.Println("OutMono: ", OutMono)
	fmt.Println("OutStereo: ", OutStereo)
	fmt.Println("OutSurround: ", OutSurround)
}

/*
OutMute:  0
OutMono:  1
OutStereo:  2
OutSurround:  5
*/
