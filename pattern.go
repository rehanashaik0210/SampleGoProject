package main
import( 
	"fmt"
)
func main(){
	x:= "REHANA"
	for i:=0; i<len(x); i++ {
		for j:=0; j<=i; j++ {
			fmt.Printf(" %c ", x[i])
		}
		fmt.Println( )
	
	}
		 
}