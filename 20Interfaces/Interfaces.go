//接口是一个方法签名的集合。
//所谓方法签名，就是指方法的声明，而不包括实现。
package _0Interfaces

import (
    "fmt"
    "math"
)

// 这里我们要让正方形square和圆形circle实现这个接口
type gemotry interface {
    area() float64
    perim() float64
}

type square struct {
    width, height float64
}

type circle struct {
    radius float64
}

// 在Go中实现一个接口，只要实现该接口定义的所有方法即可
// 下面是正方形实现的接口
func (s square) area() float64 {
    return s.width * s.height
}

func (s square) perim() float64 {
    return 2*s.width + 2*s.height
}

// 圆形实现的接口
func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}

// 如果一个函数的参数是接口类型，那么我们可以使用命名接口 来调用这个函数
// 比如这里的正方形square和圆形circle都实现了接口geometry，
// 那么它们都可以作为这个参数为geometry类型的函数的参数。
// 在measure函数内部，Go知道调用哪个结构体实现的接口方法。
func measure(g gemotry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}

func main() {
    s := square{3,4}
    c := circle{5}
    
    measure(s)
    measure(c)
}
