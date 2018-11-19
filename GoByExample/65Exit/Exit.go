// 使用os.Exit可以给定一个状态，然后立刻退出程序运行。
package _5Exit

import (
    "fmt"
    "os"
)

func main() {
    // 当使用`os.Exit`的时候defer操作不会被运行，
    // 所以这里的``fmt.Println`将不会被调用
    defer fmt.Println("!")

    // 退出程序并设置退出状态值
    os.Exit(3)
    // 注意，Go和C语言不同，main函数并不返回一个整数来表示程序的退出状态，而是将退出状态作为os.Exit函数的参数。
    // 如果你使用go run来运行程序，将会有如下输出
    // exit status 3
}
