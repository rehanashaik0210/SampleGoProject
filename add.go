package main
import"fmt"
func add(x,y int)int{
	return x+y
}
func main(){
	var num1,num2 int
	fmt.Print("enter the first number=")
	fmt.Scanln(&num1)
	fmt.Print("enter the second number=")
	fmt.Scanln(&num2)
	fmt.Println("the sum of num1 and num2=",num1+num2)
	//add(num1,num2)
}