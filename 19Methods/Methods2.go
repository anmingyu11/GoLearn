package _9Methods

import "fmt"

// 从某种意义上说，方法是函数的“语法糖”。当函数与某个特定的类型绑定，那么它就是一个方法。也证因为如此，我们可以将方法“还原”成函数。
// instance.method(args)->(type).func(instance,args)

// 为了区别这两种方式，官方文档中将 左边的称为Method Value，右边则是Method Expression。Method Value是包装后的状态对象，
// 总是与特定的对象实例关联在一起（类似闭包，拐带私奔），而Method Expression函数将Receiver作为第一个显式参数，调用时需额外传递。
// 注意：对于Method Expression，T仅拥有T Receiver方法，*T拥有（T+*T）所有方法。
type Person struct {
    Id   int
    Name string
}

func (this Person) testMethod(x int) {
    fmt.Println("id : ", this.Id, "Name : ", this.Name)
    fmt.Println("x : ", x)
}

func main() {
    p := Person{2, "zhangsan"}
    p.testMethod(1)
    var f1 func(int) = p.test
    f1(2)
    var f2 func(Person, int) = Person.test
    f2(p, 4)
}
