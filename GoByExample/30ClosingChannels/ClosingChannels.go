package _0ClosingChannels

import "fmt"

// 关闭通道的意思是该通道将不再允许写入数据。这个方法可以让通道数据的接受端知道数据已经全部发送完成了。
// 在这个例子中，我们使用通道jobs在main函数所在的协程和一个数据
// 接收端所在的协程通信。当我们数据发送完成后，我们关闭jobs通道
func main() {
    jobs := make(chan int, 5)
    done := make(chan bool)

    // 这里是数据接收端协程，它重复使用`j, more := <-jobs`来从通道
    // jobs获取数据，这里的more在通道关闭且通道中不再有数据可以接收的
    // 时候为false，我们通过判断more来决定所有的数据是否已经接收完成。
    // 如果所有数据接收完成，那么向done通道写入true
    go func() {
        for {
            if j, more := <-jobs; more {
                fmt.Println("received job", j)
            } else {
                fmt.Println("received all jobs")
                done <- true
                return
            }
        }
    }()

    // 这里向jobs通道写入三个数据，然后关闭通道
    for j := 0; j < 3; j++ {
        jobs <- j
        fmt.Println("sent job", j)
    }
    close(jobs)
    fmt.Println(" sent all jobs")

    <-done
}
