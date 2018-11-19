### Go语言有些不同的地方。

1. `:=` 初始化并赋值。
2. `const` 可以执行任意精度的数学计算
3. 常量没有确定类型，除非进行显式转换
4. 常亮可以在程序的逻辑上下文中获取类型
5. Go没有三目运算符
6. `if`可以省略小括号但是不可以省略大括号
7. `switch`语句的`case`语句块内不需要加`break`
8.  在一个 `case` 语句中，你可以使用逗号来分隔多个表达式。
9. `switch`可以不带根表达式，就成了`if/else`逻辑的另一种方式
10. `case`可以使用非常量
11. 内置函数`len()`返回数组的长度
12. `i++`和`i--`在Go语言中是语句，不是表达式，因此不能赋值给另外的变量。此外没有`++i`和`--i`。
13. Go交换两个数的顺序`s[i],s[j] = s[j],s[i]`
