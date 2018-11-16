package _4Constants

import (
	"fmt"
	"math"
)

//Go支持定义字符常量，字符串常量，布尔型常量和数值常量。
//使用const关键字来定义常量。
const s string = "constant";

func main() {
	fmt.Println(s);
	// "const"关键字可以出现在任何"var"关键字出现的地方
	// 区别是常量必须有初始值
	const n = 50000;
	// 常量表达式可以执行任意精度数学计算
	const d = 3e20 / n;
	fmt.Println(d);
    //数值型常量是没有确定的类型的，直到它们被给定了一个类型，比如说一次显示的类型转化。
	fmt.Println(int64(d))

	// 数值型常量可以在程序的逻辑上下文中获取类型
	// 比如变量赋值或者函数调用。
	// 例如，对于math包中的Sin函数,它需要一个float64类型的变量
	fmt.Println(math.Sin(n))
}
