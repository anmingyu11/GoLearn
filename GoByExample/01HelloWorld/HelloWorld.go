package _1HelloWorld

import "fmt"

//我们的第一个程序将打印传说中的 "hello world"消息，右边是完整的程序代码。
//要运行这个程序，将这些代码放到 hello-world.go 中并且使用 go run 命令。
//有时候我们想将我们的程序编译成二进制文件。我们可以通过 go build 命来达到目的。
//然后我们可以直接运行这个二进制文件。
//现在我们可以运行和并以基本的 Go 程序，让我们学习更多的关于这门语言的知识吧。
func main() {
    fmt.Println("hello world")
}