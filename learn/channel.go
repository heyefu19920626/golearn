package main

import (
    "fmt"
    "time"
)

func say(name string, index int){
    for i := 0; i < index; i++ {
        time.Sleep(100)
        fmt.Printf("Hello,%s\n", name)
    }
}

func main(){
    // go say("Li", 10)
    // go say("Mu", 10)
    // time.Sleep(10 * time.Second)
    // say("over", 10)

    s := []int{7, 2, 8, -9, 4, 0}
    c := make(chan int)
    go sum(s[:len(s)/2], c)
    time.Sleep(time.Second)
    go sum(s[len(s)/2:], c)
    x, y := <- c, <- c
    fmt.Println(x, y , x + y)

    ch := make(chan int, 2)
    ch <- 1
    ch <- 2
    fmt.Println(<- ch)
    fmt.Println(<- ch)

    ch = make(chan int, 10)
    go fibonacci(cap(ch), ch)

    for i := 0; i < cap(ch); i++ {
        v, ok := <- ch
        fmt.Println(ok)
        fmt.Println(v)
    }

    
    // for i := range ch{
    //     fmt.Println(ok, i)
    // }
    // f := func(x, y int){
    //     return x + y
    // }
    // fmt.Println(f(3,4))
}

func sum(s []int, c chan int){
    sum := 0
    for _, v := range s{
        sum += v
    }
    c <- sum
}

func fibonacci(n int, ch chan int){
    x, y := 0, 1
    for i := 0; i < n; i++ {
        ch <- x
        x, y = y, x + y
    }
    // 关闭通道并不会丢失里面的数据，只是让读取通道数据的时候不会读完之后一直阻塞等待新数据写入
    close(ch)
}


