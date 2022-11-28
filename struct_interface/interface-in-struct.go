package main

import (
	"fmt"
	"strconv"
	"time"
)

type OpenCloser interface {
	Open() error
	Close() error
}

type Door struct {
	open bool
	lock bool
}

func (d *Door) Open() error {
	fmt.Println("door open...")
	d.open = true
	return nil
}

func (d *Door) Close() error {
	fmt.Println("door close...")
	d.open = false
	return nil
}

type AutoDoor struct {
	OpenCloser
	delay int
	msg   string
}

func (a *AutoDoor) Open() error {
	fmt.Println("Open after " + strconv.Itoa(a.delay) + " seconds")
	time.Sleep(time.Duration(a.delay) * time.Second)
	fmt.Println("Door is opening:" + a.msg)
	return nil
}

func main() {
	door := &AutoDoor{&Door{false, false}, 3, "warning"}
	door.Open()
	if v, ok := door.OpenCloser.(*Door); ok {
		fmt.Println(v)
	}

	door.OpenCloser.Open()
	if v, ok := door.OpenCloser.(*Door); ok {
		fmt.Println(v)
	}

	door.Close()
	if v, ok := door.OpenCloser.(*Door); ok {
		fmt.Println(v)
	}

}

/*
$ go run interface-in-struct.go
Open after 3 seconds
Door is opening:warning
&{false false}
door open...
&{true false}
door close...
&{false false}
*/
