package main 

import "fmt"

func main() {
    //maps are go's associative data type (like a dict or hash)
    map1 := make(map[string]int)
    map1["a"] = 1
    map1["b"] = 2

    fmt.Println(map1)

    //len function counts k:v pairs
    fmt.Println("len: ",len(map1))

    //builtin delete removes k:v pairs given a map and key
    delete(map1, "a")
    fmt.Println(map1)

    //accessing a value given a key has a second optional return value:
    //a bool indicating whether the key exists
    val, prs := map1["a"]
    fmt.Println(val, prs)
    //since an absent key will map to 0 by default, we need this second return
    //to distinguish between a key mapped to zero vs. a key that doesn't exist

    //maps can also be initialized in one line with a map literal
    map2 := map[string]int{"foo":1, "bar":2}
    fmt.Println(map2)

}