// 命令行参数标记是为命令行程序指定选项参数的常用方法。
// 例如，在命令wc -l中，-l就是一个命令行参数标记。
// Go提供了flag包来支持基本的命令行标记解析。
// 我们这里将要使用这个包提供的方法来实现带选项的命令行程序。
package _0CommandLineFlags

import (
    "flag"
    "fmt"
)

func main() {
    // 基础的标记声明适用于string，integer和bool型选项。
    // 这里我们定义了一个标记`word`，默认值为`foo`和一个简短的描述。
    // `flag.String`函数返回一个字符串指针（而不是一个字符串值），我们下面将演示如何使用这个指针
    wordPtr := flag.String("word", "foo", "a string")

    // 这里定义了两个标记，一个`numb`，另一个是`fork`，
    // 使用和上面定义`word`标记相似的方法
    numbPtr := flag.Int("numb", 42, "an int")
    boolPtr := flag.Bool("fork", false, "a bool")

    // 你也可以程序中任意地方定义的变量来定义选项，只
    // 需要把该变量的地址传递给flag声明函数即可
    var svar string
    flag.StringVar(&svar, "svar", "bar", "a string var")

    // 当所有的flag声明完成后，使用`flag.Parse()`来分
    // 解命令行选项
    flag.Parse()

    // 这里我们仅仅输出解析后的选项和任何紧跟着的位置
    // 参数，注意我们需要使用`*wordPtr`的方式来获取最
    // 后的选项值
    fmt.Println("word:", *wordPtr)
    fmt.Println("numb:", *numbPtr)
    fmt.Println("fork:", *boolPtr)
    fmt.Println("svar:", svar)
    fmt.Println("tail:", flag.Args())
}
