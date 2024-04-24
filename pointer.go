package main

import "fmt"

func main() {
	a := 77
	fmt.Println(a)
	fmt.Println(&a)
	b := &a
	fmt.Println(b)
	fmt.Println(*b)
	fmt.Println(*&b)
	*b = 47
	fmt.Println(a)
}
