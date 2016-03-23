package main 

import "fmt"

func main() {

    //initialize a zero-valued integer array of length 5
    var a [5]int
    fmt.Println(a)

    //2D nested array
    var twoD [2][3]int
    for i:=0; i<2; i++ {
        for j:=0; j<3; j++{
            twoD[i][j]=i+j
        }
    }
    fmt.Println("2d: ",twoD)

    //slices are a more powerful alternative to arrays in Go
    s:=make([]string,3)
    s[0]="a"
    s[1]="b"
    s[2]="c"
    fmt.Println(s)
    fmt.Println(s[1])
    fmt.Println("len: ",len(s))
    //slices are variable length
    s = append(s, "d")
    s = append(s, "e", "f")
    fmt.Println("len: ",len(s))

    // a slice can be copied to another of the same length
    s2:=make([]string,6)
    copy(s2, s)
    fmt.Println(s, s2)

    //a slice literal can be delcared without specifying length
    letters := []string{"a","b","c","d"}
    fmt.Println(letters)

    //a slice can be created by slicing another slice or array
    sub_letters := letters[1:3]
    fmt.Println(sub_letters)

    //for 2D slices, inner dimension can vary in size!
    twoDSlice := make([][]int, 3)
    for i:=0; i < 3; i++ {
        twoDSlice[i] = make([]int, i+1)
        for j:=0; j<i+1; j++ {
            twoDSlice[i][j] = j + i
        }
    }
    fmt.Println(twoDSlice)


}