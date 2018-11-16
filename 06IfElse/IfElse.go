package _6IfElse

import "fmt"

//条件判断结构中，条件两边的小括号()是可以省略的，但是条件执行语句块两边的大括号{}不可以。
//Go 里没有三目运算符，所以即使你只需要基本的条件判断，你仍需要使用完整的 if 语句。
func main() {
    // 基本的例子 
    if 7%2 == 0 {
        fmt.Println(" 7 is even ")
    } else {
        fmt.Println(" 7 is odd ")
    }

    // 只有if条件的情况
    if 8%4 == 0 {
        fmt.Println(" 8 is divided by 4")
    }

    // if条件可以包含一个初始化表达式，这个表达式中的变量
    // 是这个条件判断结构的局部变量
    if num := 9; num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }
}
