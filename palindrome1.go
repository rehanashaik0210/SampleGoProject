package main
import"fmt"
func main(){
	var palNum,remainder int
	fmt.Print("Enter the Number to check palindrome=")
	fmt.Scanln(&palNum)
	reverse:=0
	for temp:=palNum;temp>0;temp=temp/10{
		remainder=temp%10
		reverse=reverse*10+remainder
	}
	fmt.Println("The reverse of the given number=",reverse)
	if palNum==reverse{
		fmt.Println(palNum,"is a palindrome number")
	}else{
		fmt.Println(palNum,"is not a palindrome number")
	}
}