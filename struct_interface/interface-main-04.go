package main

import (
	"log"
)

type Notifier interface {
	Notify() error
}

func SendNotification(notify Notifier) error {
	return notify.Notify()
}

type User struct {
	Name  string
	Email string
}

func (u *User) Notify() error {
	log.Printf("User: Sending User Email To %s<%s>\n", u.Name, u.Email)
	return nil
}

func main() {
	user := User{
		Name:  "AriesDevil",
		Email: "ariesdevil@xxoo.com",
	}

	SendNotification(&user)
}

/*
类型 T 的可调用方法集包含接受者为 T 的所有方法
这条规则说的是如果我们用来调用特定接口方法的接口变量是一个值类型，那么方法的接受者必须也是值类型该方法才可以被调用。

类型 *T 的可调用方法集包含接受者为 *T 或 T 的所有方法集
这条规则说的是如果我们用来调用特定接口方法的接口变量是一个指针类型，那么方法的接受者可以是值类型也可以是指针类型。

类型 T 的可调用方法集不包含接受者为 *T 的方法


$ go run interface-main-04.go
2022/11/28 13:13:10 User: Sending User Email To AriesDevil<ariesdevil@xxoo.com>

*/
