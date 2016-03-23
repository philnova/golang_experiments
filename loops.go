package main

import "fmt"

func main() {
    //For is the only loop construct in go

    i:=1
    for i<=3{
        fmt.Println(i)
        i++
    }

    //alternatively a For loop can be used with more C-like syntax
    for j:=7; j<=9; j++ {
        fmt.Println(j)
    }

    i = 0
    //a For loop without a condition loops infinitely until break or return
    for {
        fmt.Println("loop")
        i++
        if i==3 {
            break
        }
        
    }
}