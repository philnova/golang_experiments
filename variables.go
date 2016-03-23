package  main

import "fmt"
import "math"

func main() {
    //variables are declared as var variable_name type
    var a string = "initial"
    fmt.Println(a)

    //multiple variables can be declared on the same line
    var b, c int = 1, 2
    fmt.Println(b)
    fmt.Println(c)

    //:= syntax is shorthand for declaring and initializing a variable
    //type is inferred by the compiler if not stated
    f:="short"
    fmt.Println(f)

    //const can replace var to declare a constant
    const s string = "constant"

    //numerical constants have no type until they are given one
    const n = 50

    //for example, math.Sin expects a float64, so n becomes that type
    fmt.Println(math.Sin(n))



}