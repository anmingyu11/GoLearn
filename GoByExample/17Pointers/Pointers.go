package _7Pointers

import "fmt"

//Go 支持 指针，允许在程序中通过引用传递值或者数据结构
func main() {
    i := 1
    fmt.Println("initial : ", i)

    zeroval(i)
    fmt.Println("zeroval:", i)

    // &操作符用来取得i变量的地址
    zeroptr(&i)
    fmt.Println("zeroptr", i)

    // 指针类型也可以输出
    fmt.Println("pointer:", &i)
    
}

// 我们将通过两个函数：zeroval 和 zeroptr 来比较指针和值类型的不同。
// zeroval 有一个 int 型参数，所以使用值传递。
// zeroval 将从调用它的那个函数中得到一个 ival形参的拷贝。

// 我们用两个不同的例子来演示指针的用法
// zeroval函数有一个int类型参数，这个时候传递给函数的是变量的值
func zeroval(ival int) {
    ival = 0
}

// zeroptr函数的参数是int类型指针，这个时候传递给函数的是变量的地址
// 在函数内部对这个地址所指向的变量的任何修改都会反映到原来的变量上。
func zeroptr(iptr *int) {
    *iptr = 0
}
