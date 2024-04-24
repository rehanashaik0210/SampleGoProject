package main
import"fmt"
func main(){
	var num,count int
	count=0
	fmt.Print("Enter the number to count digits=")
	fmt.Scanln(&num)
	for num>0{
		num=num/10
		count=count+1
	}
	fmt.Println("The total number to digits=",count)
}