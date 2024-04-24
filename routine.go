package main
import(
	"fmt"
)
func main(){
	go func(){
		fmt.Println("print the first thing")
	}()
	go func(){
		fmt.Println("print the second thing")
	}()
	fmt.Println("about to exit")
}