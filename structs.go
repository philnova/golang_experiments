package main 

import "fmt"

//structs allow us to define custom composite types with fields containing other types
type person struct {
    name string
    age int
}

func main() {

    //create a struct literal with {}
    fmt.Println(person{"Bob", 20})

    //can use named fields when initializing a struct
    sara := person{name : "Sara", age : 20}
    fmt.Println(sara)

    //when using named fields, the order need not match the struct definition
    sara = person{age : 24, name : "Sara"}
    fmt.Println(sara)

    //structs are mutable
    sara.age = 25
    fmt.Println(sara)

    //access struct fields with . syntax
    sean := person{name: "Sean", age: 50}
    fmt.Println(sean.name)

    //point to structs with standard syntax
    sp := &sean
    fmt.Println(sp.age)

    //can also use . syntax on pointers to structs
    //the pointer is automatically dereferenced
    sp.age = 51
    fmt.Println(sp.age)
}