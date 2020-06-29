package main

import (
    "fmt"
    "os"
    "math/rand"
    "sync"
)

//生成随机数
func produce(data chan int, wg *sync.WaitGroup) {
    n := rand.Intn(999)
    data <- n
    wg.Done()
}
//写入文件
func consume(data chan int, done chan bool) {
    f, err := os.Create("concurrent")
    if err != nil {
        fmt.Println(err)
        return
    }
    for d := range data {
        _, err = fmt.Fprintln(f, d)
        if err != nil {
            fmt.Println(err)
            f.Close()
            done <- false
            return
        }
    }
    err = f.Close()
    if err != nil {
        fmt.Println(err)
        done <- false
        return
    }
    done <- true
}

//启用100个goroutines生成随机数并写入
func main() {
    data := make(chan int)
    done := make(chan bool)
    wg := sync.WaitGroup{}
    for i := 0; i < 100; i++ {
        wg.Add(1)
        go produce(data, &wg)
    }
    go consume(data, done)
    go func() {
        wg.Wait()
        close(data)
    }()
    d := <-done
    if d == true {
        fmt.Println("File written successfully")
    } else {
        fmt.Println("File written successfully")
    }
}
