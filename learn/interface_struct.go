package main

import "fmt"


/**
 * 定义学生接口
 */
type Student interface{
    name() string
    age() int
}

/**
 * 定义男孩结构体
 */
type Boy struct{

}

/**
 * 男孩结构体实现学生接口的name方法
 */
func (boy Boy) name() string{
    return "boy name"
}

func (boy Boy) age() int{
    return 19
}

type Girl struct{

}

func (girl Girl) name() string{
    return "girl name"
}

func (girl Girl) age() int{
    return 17
}

func main(){
    var stu Student

    stu = new(Boy)
    fmt.Println(stu.name())
    fmt.Println(stu.age())

    stu = new(Girl)
    fmt.Println(stu.name())
    fmt.Println(stu.age())
}