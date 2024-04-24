package main

import "fmt"

func main() {
	var num int 
	fmt.Println("enter the value to find cube")
	fmt.Scanln(&num)
	cube := num * num * num
	fmt.Println("\nthe cube of a given num is=", cube)
}