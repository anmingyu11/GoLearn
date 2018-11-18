//读写文件是很多程序的基本任务，下面我们看看Go里面的文件读取。
package _6ReadingFile

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 读取文件的函数调用大多数都需要检查错误，
// 使用下面这个错误检查方法可以方便一点
func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	const fileName = "/home/amy/dump/atomic"
	// 最基本的文件读写任务就是把整个文件的内容读取到内存
	dat, err := ioutil.ReadFile(fileName)
	check(err)
	fmt.Println(string(dat))

	// 有的时候你想更多地控制到底是读取文件的哪个部分，这个
	// 时候你可以使用`os.Open`打开一个文件获取一个`os.File`
	// 对象

	f, _ := os.Open(fileName)

	// 从这个文件中读取一些字节，并且由于字节数组长度所限，
	// 最多读取5个字节，另外还需要注意实际能够读取的字节数量
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	// 你也可以使用`Seek`来跳转到文件中的一个已知位置，并从
	// 那个位置开始读取数据
	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))

	// `io`包提供了一些帮助文件读取的函数。例如上面的方法如果
	// 使用方法`ReadAtLeast`函数来实现，将使得程序更健壮
	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 10)
	n3, err := io.ReadAtLeast(f, b3, 2)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	// `bufio`包提供了缓冲读取文件的方法，这将使得文件读取更加高效
	r4 :=bufio.NewReader(f)
	b4,err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

}
