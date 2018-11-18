package _1RangeOverChannels

import "fmt"

//我们知道range函数可以遍历数组，切片，字典等。这里我们还可以使用range函数来遍历通道以接收通道数据。
func main() {
    // 我们遍历queue通道里面的两个数据
    queue := make(chan int, 2)
    queue <- 1
    queue <- 2
    close(queue)

    // range函数遍历每个从通道接收到的数据，因为queue再发送完两个
    // 数据之后就关闭了通道，所以这里我们range函数在接收到两个数据
    // 之后就结束了。如果上面的queue通道不关闭，那么range函数就不
    // 会结束，从而在接收第三个数据的时候就阻塞了。
    for element := range queue {
        fmt.Println(element)
    }
}
