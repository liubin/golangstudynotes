package main

import "fmt"

func main() {

    var arr = [5]int {1,2,3,4,5}

    for i :=0 ; i < len(arr) ; i++ {
        fmt.Println("element ", i, " is ", arr[i])
    }

    fmt.Println("")

    for i,v := range arr {
        fmt.Println("element ", i, " is ", v)
    }

    fmt.Println("")

    // arr2 is an array slice type
    // Array Slice case 1: create from array
    var arr2 []int = arr[2:4]
    for _,v := range arr2 {
        fmt.Println(v)
    }
    // TODO codes above print 3 4 ,not include 2

    fmt.Println("")

    // Array Slice case 2: create a array without any elements.
    var arr3 []int
    // with all elements is set to zeor
    arr3 = make([]int, 3)
    for _,v := range arr3 {
        // will print 0
        fmt.Println(v)
    }

    // or
    arr4 := []int {1,2,3,4,5}

    fmt.Println("")

    // append element(s) to array slice
    // with element directly
    arr4 = append(arr4, 6,7,8)
    
    // please pay attention to arr3...
    // this will sperate elements from arr3 and as an int... parameter list
    arr4 = append(arr4, arr3...)

    for _,v := range arr4 {
        fmt.Println(v)
    }


}
