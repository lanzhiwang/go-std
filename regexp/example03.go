package main

import (
	"fmt"
	"regexp"
)

func main() {
	// func MustCompile(str string) *Regexp
	re := regexp.MustCompile("^[a-zA-Z0-9]{5,12}")

	// func (re *Regexp) MatchString(s string) bool
	fmt.Println(re.MatchString("slimshady99"))
	fmt.Println(re.MatchString("!asdf£33£3"))
	fmt.Println(re.MatchString("roger"))
	fmt.Println(re.MatchString("iamthebestuserofthisappevaaaar"))
}
