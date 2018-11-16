package _3MultipleReturnValues

import "fmt"

// Go语言内置支持多返回值，这个在Go语言中用的很多，比如一个函数同时返回结果和错误信息。
func main() {
    // 获取函数的两个返回值
    a, b := vals()
    fmt.Println(a)
    fmt.Println(b)

    // 如果你只对多个返回值里面的几个感兴趣
    // 可以使用下划线(_)来忽略其他的返回值
    _, c := vals()
    fmt.Println(c)
}

func vals() (int, int) {
    return 3, 7;
}
