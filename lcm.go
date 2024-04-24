package main
import(
	"fmt"
)
func gcd(a,b int) int{
	if b==0{
		return a
	}
	return gcd (b,a%b)
}
func lcm (i,j int) int{
	return (i*j)/gcd(i,j)

}
func main(){
	var a,b int
	fmt.Scanln(&a,&b)
	fmt.Println("The lcm of %d and %d is %d:" ,a,b,lcm(a,b))
		
	}
