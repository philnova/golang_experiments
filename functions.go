package main 

import "fmt"

//example function declaration
//declare types of arguments and return value
func sum(a int, b int) int {
    return a + b
}

//when arguments have the same type, you may omit type declaration until the end
func sum3(a, b, c int) int {
    return a + b + c
}

//a function may return multiple arguments
//simply specify a type for each argument
func tautology(a, b int) (int, int) {
    return a, b
}

//variadic functions (like fmt.Println) can accept arbitrary numbers of arguments
func sum_arbitrary(nums ...int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}

//go supports closures
//closures are functions that modify local variables defined outside their own scope
func intSeq() func() int {
    i := 0
    return func() int {
        i+=1
        return i
    }
}

//go supports recursion
func fact(n int) int {
    if n == 0 {
        return 1
    }
    return n * fact(n-1)
}

func main() {
    res := sum(1,2)
    fmt.Println("1+2=",res)

    res = sum3(1,2,3)
    fmt.Println("1+2+3=",res)

    v1, v2 := tautology(1,2)
    fmt.Println("1,2 =",v1, v2)

    fmt.Println(sum_arbitrary(1,2,3,4,5,6))

    //a variadic function can be applied to a slice using func(slice...) syntax
    nums := []int{1,2,3,4}
    fmt.Println(sum_arbitrary(nums...))

    //closures allow us to replicate the generator pattern as in Python
    nextInt := intSeq()
    for j:=0; j<10; j++ {
        fmt.Println("intSeq: ",nextInt())
    }

    fmt.Println("7!=",fact(7))
    
}