// 环境变量是一种很普遍的将配置信息传递给Unix程序的机制。
package _1EnvironmentVariables

import (
    "fmt"
    "os"
)

func main() {
    // 为了设置一个key/value对，使用`os.Setenv`
    // 为了获取一个key的value，使用`os.Getenv`
    // 如果所提供的key在环境变量中没有对应的value，
    // 那么返回空字符串
    os.Setenv("FOO", "1")
    fmt.Println(os.Getenv("FOO"))
    fmt.Println(os.Getenv("BAR"))

    // 使用`os.Environ`来列出环境变量中所有的key/value对
    // 你可以使用`strings.Split`方法来将key和value分开这里我们打印所有的key
    for _, e := range os.Environ() {
        fmt.Println(e)
    }

    // 这里我们设置了FOO环境变量，所以我们取到了它的值，但是没有设置BAR环境变量，所以值为空。
    // 另外我们列出了系统的所有环境变量，当然这个输出根据不同的系统设置可能并不相同。
}
