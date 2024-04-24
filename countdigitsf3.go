package main
import"fmt"
func digitCount(num int)int{
	var count int=0
	for num>0{
		num=num/10
		count=count+1
	}
	return count
}
func main(){
	var num ,count int
	fmt.Print("Enter any number to count digits=")
	fmt.Scanln(&num)
	count=digitCount(num)
	fmt.Println("The total number to digits=",count)
}