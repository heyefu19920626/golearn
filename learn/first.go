package main

import "fmt"

const (
    a = iota //0
    b
    c
    d = "haha"
    e = iota //4
    f
)

// 在下一个 const 出现之前，每出现一次 iota,其所代表的数字自动加 1
const g = iota //0

func main(){
    fmt.Println(a, b, c, d, e, f, g)
    fmt.Println("Hello Go!")

    var a = 1
    a, b := 2, 3

    fmt.Println(a, b)

    // 空白标识符用于抛弃值
    _, c , d := numbers()
    fmt.Println(c, d)

    // 常量
    const LENGTH  = 9

    // 运算符
    // &,|,^,<<,>> 双目
    // &,*单目
    ptr := &a
    fmt.Println(*ptr)
    // 切片,限定长度的为数组,[...]
    var array = [...]int {1,2,3}
    fmt.Println(array)
}

func numbers()(int, int, int){
    a, b, c := 1, 2, 3

    return a, b, c
}