package main 

import "fmt"

func main() {
    nums := []int{2,3,4}
    sum := 0

    //range provides both indices and values, like Python's enumerate()
    //we may not always need the indices
    for _, num := range nums {
        sum += num
    }
    fmt.Println("sum: ", sum)

    for i, num := range nums {
        if num == 3 {
            fmt.Println("index:", i)
        }
    }

    kvs := map[string]string{"a": "apple", "b":
    "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }

    //ranges on strings iterate over Unicode code points
    for i, c := range "go" {
        fmt.Println(i, c)
    }
}