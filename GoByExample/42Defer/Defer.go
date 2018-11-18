// Defer 被用来确保一个函数调用在程序执行结束前执行。同样用来执行一些清理工作。 
// defer 用在像其他语言中的ensure 和 finally用到的地方。
package _2Defer

import (
    "fmt"
    "os"
)

// 假设我们想创建一个文件，然后写入数据，最后关闭文件
func main() {
    // 在使用createFile得到一个文件对象之后，我们使用defer
    // 来调用关闭文件的方法closeFile，这个方法将在main函数
    // 最后被执行，也就是writeFile完成之后

    f := createFile("/home/amy/dump/dumpGo")

    defer closeFile(f)
    writeFile(f)
}

func createFile(p string) *os.File {
    fmt.Println("creating")
    f, err := os.Create(p);
    if err != nil {
        panic(err)
    }
    return f
}

func writeFile(file *os.File) {
    fmt.Println("writing")
    fmt.Fprintf(file, "fuck")
}

func closeFile(file *os.File) {
    fmt.Println("closing")
    file.Close();
}
