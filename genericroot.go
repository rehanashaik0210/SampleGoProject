package main
import"fmt"
func main(){
	var gennumber,sum,remainder int
	fmt.Print("enter the number to find generic root=")
	fmt.Scanln(&gennumber)
	for gennumber>=10{
		for sum=0;gennumber>0;gennumber=gennumber/10{
			remainder=gennumber%10
			sum=sum+remainder
		}
		if sum>=10{
			gennumber=sum
		}else{
			 		}
	}
	fmt.Println("The generic root of this number=",sum)

}