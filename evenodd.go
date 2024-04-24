package main

import "fmt"

func main() {

    var eonum int

    fmt.Print("Enter the Number to find Even or Odd = ")
    fmt.Scanln(&eonum)

    if eonum%2 == 0 {
        fmt.Println("\nIt is an Even Number")
    } else {
        fmt.Println("\nIt is an Odd Number")
    }
}