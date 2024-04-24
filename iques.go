package main

import "fmt"

//type Student struct {
//	Name       string
//	RollNumber int
//}

func main() {
	//		student := Student{
	//			Name:       "Rifa",
	//			RollNumber: 1,
	//		}
	//		fmt.Println("Student Details:")
	//		fmt.Println("Name:", student.Name)
	//		fmt.Println("RollNum:", student.RollNumber)
	//	}
	var numbers [3]int
	numbers[0] = 34
	numbers[1] = 1
	numbers[2] = 5
	fmt.Println("Original Array:")
	fmt.Println("Element at index 0:", numbers[0])
	fmt.Println("Element at index 1:", numbers[1])
	fmt.Println("Element at index 2:", numbers[2])
}
