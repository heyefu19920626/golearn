package main

/**
* 流程控制
*/

import(
    "fmt"
    "errors"
    // "math"
)

func main(){
    fmt.Println("Hello Go!")
    x := -2
    test_goto(x)
    classchecker(5, -17.98, "AIDEN", nil, true, complex(1, 1))
    // fmt.Println(math.add(4,5))
    // devide(4, 2)
    fmt.Println(devide(4, 2))
    result, err := devide(4, 0)
    if err != nil{
        fmt.Println(err)
    }else{
        fmt.Println(result)
    }

    // 匿名函数
    y := func(a int, b int, c int) bool{
        return a + b < c
    }(1, 3, 2)
    fmt.Println(y)

    // 类型断言
    type_accest()

    // defer
    fmt.Println(test_defer())

    panic_recover()
}

func panic_recover()(x int){
    defer func(){
        if e := recover(); e != nil{
            err := e.(error)
            fmt.Println("Catch ecpection: ", err)
        }
    }()
    panic(errors.New("bug"))
    fmt.Println("----")
    return 10
}

func test_defer()(x int){
    // defer 在return之后执行
    // defer 可以使用匿名函数，匿名函数必须自调用
    defer func (){
        fmt.Println("defer: 20")
    }()
    return 10
}


func type_accest(){
    x := uint16(65000)
    y := int16(x) // 将 x转换为int16类型
    fmt.Printf("type and value of x is: %T and %d\n", x, x) // %T 格式化指令的作用是输出变量的类型
    fmt.Printf("type and value of y is: %T and %d\n", y, y)

    var i interface{} = 99 // 创建一个interface{}类型，其值为99
    var s interface{} = []string{"left", "right"}
    j := i.(int) // 我们假设i是兼容int类型，并使用类型断言将其转换为int类型
    fmt.Printf("type and value of j is: %T and %d\n", j, j)

    if s, ok := s.([]string); ok { // 创建了影子变量，if的作用域中覆盖了外部的变量s
        fmt.Printf("%T -> %q\n", s, s)
    }
}

func devide(a int, b int)(num int, err error){
    if b == 0{
        err := errors.New("除数不能为0！")
        return 0, err
    }
    return a / b, nil
}

func test_goto(x int){
    THIS:
    fmt.Println("THIS x < 0")
    x++
    if x < 0 {
        goto THIS
    }
}

func classchecker(items ...interface{}) { // 创建一个函数，该函数可以接受任意多的任意类型的参数
    for i, x := range items {
        switch x := x.(type) { // 创建了影子变量
        case bool:
            fmt.Printf("param #%d is a bool, value: %t\n", i, x)
        case float64:
            fmt.Printf("param #%d is a float64, value: %f\n", i, x)
        case int, int8, int16, int32, int64:
            fmt.Printf("param #%d is a int, value: %d\n", i, x)
        case uint, uint8, uint16, uint32, uint64:
            fmt.Printf("param #%d is a uint, value: %d\n", i, x)
        case nil:
            fmt.Printf("param #%d is a nil\n", i)
        case string:
            fmt.Printf("param #%d is a string, value: %s\n", i, x)
        default:
            fmt.Printf("param #%d's type is unknow\n", i)
        }
    }
}