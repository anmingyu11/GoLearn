// 我们经常需要程序去处理一些集合数据，比如选出所有符合条件的数据或者使用一个自定义函数将一个集合元素拷贝到另外一个集合。
// 在一些语言里面，通常是使用泛化数据结构或者算法。
// 但是Go不支持泛化类型，在Go里面如果你的程序或者数据类型需要操作集合， 那么通常是为集合提供一些操作函数。
// 这里演示了一些操作strings切片的集合函数，你可以使用这些例子来构建你自己的函数。
// 注意在有些情况下，使用内联集合操作代码会更清晰，而不是去创建新的帮助函数。
package _3CollectionFunctions

import (
    "fmt"
    "strings"
)

func Index(vs []string, t string) int {
    for i, s := range vs {
        if t == s {
            return i;
        }
    }
    return -1;
}

// 如果t存在于vs中，那么返回true，否则false
func Include(vs []string, t string) bool {
    return Index(vs, t) >= 0
}

func Any(vs []string, f func(string) bool) bool {
    for _, v := range vs {
        if (f(v)) {
            return true
        }
    }
    return false
}

func All(vs []string, f func(string) bool) bool {
    for _, v := range vs {
        if !f(v) {
            return false
        }
    }
    return true
}

// 返回一个新的字符串切片，切片的元素为vs中所有能够让函数f
// 返回true的元素
func Filter(vs []string, f func(string) bool) []string {
    newVs := make([]string, 0)
    for _, v := range vs {
        if f(v) {
            newVs = append(newVs, v)
        }
    }

    return newVs
}

//返回一个对原始切片中所有字符串执行函数 f 后的新切片。
func Map(vs []string, f func(string) string) []string {
    newVs := make([]string, len(vs))
    for i, v := range vs {
        newVs[i] = f(v)
    }
    return newVs
}

func main() {
    var strs = []string{"peach", "apple", "pear", "plum"}
    fmt.Println(Index(strs, "pear"))
    fmt.Println(Include(strs, "grape"))

    fmt.Println(Any(strs, func(s string) bool {
        return strings.HasPrefix(s, "p");
    }))
    fmt.Println(All(strs, func(s string) bool {
        return strings.HasPrefix(s, "p")
    }))
    fmt.Println(Filter(strs, func(s string) bool {
        return strings.Contains(s, "e")
    }))

    fmt.Println(Map(strs, strings.ToUpper))
}
