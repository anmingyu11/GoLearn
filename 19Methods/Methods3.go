package _9Methods

import "fmt"

type Student struct {
    Person
    Score int
}

func (this Person) testMethod2() {
    fmt.Println("person test")
}

func (this Student) testMethod2() {
    fmt.Println("student test")
}

// 使用匿名字段，实现模拟继承。即可直接访问匿名字段（匿名类型或匿名指针类型）的方法这种行为类似“继承”。
// 访问匿名字段方法时，有隐藏规则，这样我们可以实现override效果。
func main() {
    p := Student{Person{2, "zxhang"}, 25}
    p.testMethod2()
}
