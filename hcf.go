package main
import"fmt"
func main(){
	var x,y int
	fmt.Print("Enter the two numbers:")
	fmt.Scanln(&x,&y)
	hcf:=FindHCF(x,y)
	fmt.Printf("The higest common factor of %d and %d is %d",x,y,hcf)
}
func FindHCF(a,b int) int{
	for b!=0{
		temp:=b
	b=a%b
	a=temp
}
return a
}