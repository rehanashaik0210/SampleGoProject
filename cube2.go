package main

import "fmt"

func find_cube(x int) int {
    return x * x * x
}

func main() {

    var num int

    fmt.Print("Enter the Number to find Cube = ")
    fmt.Scanln(&num)

    cube := find_cube(num)

    fmt.Println("\nThe Cube of a Given Number  = ", cube)
}