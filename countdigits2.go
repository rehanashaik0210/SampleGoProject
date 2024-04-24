package main
import"fmt"
func main(){
	var num,count int
	fmt.Print("Enter any number to count digits= ")
	fmt.Scanln(&num)
	for count=0;num>0;num=num/10{
		count=count+1
	}
	fmt.Println("the total number to digits=",count)
}
