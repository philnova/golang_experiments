package main 

import "fmt"


//zeroval receives an integer called ival
//it creates a distinct local variable of the same name and sets it to zero
//the argument will not be changed
func zeroval(ival int) {
    ival = 0
}

//zeroptr receives a pointer to an integer called iptr
//it uses the pointer to set the integer pointed to by iptr to zero
func zeroptr(iptr *int) {
    *iptr = 0
}

func main() {
    i:=1
    fmt.Println("initial i=",i)

    zeroval(i)
    fmt.Println("after zeroval(i), i=",i)

    //use &variable to receive a pointer to that variable
    zeroptr(&i)
    fmt.Println("after zeroptr(i), i=",i)

    fmt.Println("the pointer to i is at", &i)
}