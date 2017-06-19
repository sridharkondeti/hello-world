package main

import ( "fmt"
         "time"
       )

func func1(c chan int) {
    fmt.Println("In thread")
    c<-1
}

func main() {
    start := time.Now()
    c := make(chan int)
    go func1(c)
    go func1(c)
    <-c
    elapsed := time.Since(start)
    fmt.Println(elapsed)
}
