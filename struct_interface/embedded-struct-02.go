package main

/*
给定一个结构体类型 S 和一个命名为 T 的类型，方法提升像下面规定的这样被包含在结构体方法集中：

如果 S 包含一个匿名字段 T，S 和 *S 的方法集都包含接受者为 T 的方法提升。
	这条规则说的是当我们嵌入一个类型，嵌入类型的接受者为值类型的方法将被提升，可以被外部类型的值和指针调用。

	对于 *S 类型的方法集包含接受者为 *T 的方法提升
	这条规则说的是当我们嵌入一个类型，可以被外部类型的指针调用的方法集只有嵌入类型的接受者为指针类型的方法集，也就是说，当外部类型使用指针调用内部类型的方法时，只有接受者为指针类型的内部类型方法集将被提升。

如果 S 包含一个匿名字段 *T，S 和 *S 的方法集都包含接受者为 T 或者 *T 的方法提升
	这条规则说的是当我们嵌入一个类型的指针，嵌入类型的接受者为值类型或指针类型的方法将被提升，可以被外部类型的值或者指针调用。

这就是语言规范里方法提升中仅有的三条规则，根据这个推导出一条规则：
如果 S 包含一个匿名字段 T，S 的方法集不包含接受者为 *T 的方法提升。
这条规则说的是当我们嵌入一个类型，嵌入类型的接受者为指针的方法将不能被外部类型的值访问。

*/

import (
	"fmt"
)

// 命名为 T 的类型
type User struct {
	Name  string
	Email string
}

func (u User) Notify1() {
	fmt.Println("Notify1")
}

func (u *User) Notify2() {
	fmt.Println("Notify2")
}

// 结构体类型 S
type Admin struct {
	*User
	Level string
}

func main() {
	// S
	admin1 := Admin{
		User: &User{
			Name:  "AriesDevil",
			Email: "ariesdevil@xxoo.com",
		},
		Level: "super",
	}

	// *S
	admin2 := &Admin{
		User: &User{
			Name:  "AriesDevil",
			Email: "ariesdevil@xxoo.com",
		},
		Level: "super",
	}

	admin1.Notify1()
	admin1.Notify2()

	admin2.Notify1()
	admin2.Notify2()

}

/*
$ go run  embedded-struct-02.go
Notify1
Notify2
Notify1
Notify2
*/
