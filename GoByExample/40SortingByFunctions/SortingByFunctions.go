// 有的时候我们希望排序不是仅仅按照自然顺序排序。
// 例如，我们希望按照字符串的长度来对一个字符串数组排序而不是按照字母顺序来排序。这里我们介绍一下Go的自定义排序。
package _0SortingByFunctions

import (
    "fmt"
    "sort"
)

// 为了能够使用自定义函数来排序，我们需要一个
// 对应的排序类型，比如这里我们为内置的字符串
// 数组定义了一个别名ByLength
type ByLength []string

func (s ByLength) Len() int {
    return len(s)
}

func (s ByLength) Swap(i int, j int) {
    s[i], s[j] = s[j], s[i]
}

func (s ByLength) Less(i int, j int) bool {
    return len(s[i]) < len(s[j])
}

// 一切就绪之后，我们就可以把需要进行自定义排序
// 的字符串类型fruits转换为ByLength类型，然后使用
// sort包的Sort方法来排序
func main() {
    fruits := []string{"peach", "banana", "kiwi"}
    sort.Sort(ByLength(fruits))
    fmt.Println(fruits)
}
// 同样的，对于其他的类型，使用这种方法，我们可以为Go的切片提供任意的排序方法。归纳一下就是：
// 1. 创建自定义排序类型
// 2. 实现sort包的接口方法Len,Swap和Less
// 3. 使用sort.Sort方法来排序
