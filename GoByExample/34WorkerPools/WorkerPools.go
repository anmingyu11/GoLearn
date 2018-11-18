package _4WorkerPools

import (
    "fmt"
    "time"
)

// 在这个例子中，我们来看一下如何使用gorouotine和channel来实现工作池。
func main() {

    // 为了使用我们的工作池，我们需要发送工作和接受工作的结果，
    // 这里我们定义两个通道，一个jobs，一个results
    jobs := make(chan int, 100)
    results := make(chan int, 100)

    // 这里启动3个worker协程，一开始的时候worker阻塞执行，因为
    // jobs通道里面还没有工作任务
    for w := 0; w < 3; w++ {
        go worker(w, jobs, results)
    }

    // 这里我们发送9个任务，然后关闭通道，告知任务发送完成
    for j := 1; j <= 9; j++ {
        jobs <- j
    }
    close(jobs)

    // 然后我们从results里面获得结果
    for r := 1; r <= 9; r++ {
        <-results
    }
}

func worker(id int, jobs <-chan int, result chan<- int) {
    for j := range jobs {
        fmt.Println("worker ", id, " processing job ", j)
        time.Sleep(time.Millisecond * 500)
        result <- j
    }
}
