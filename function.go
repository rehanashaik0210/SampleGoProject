package main
import(
	"fmt"
)
func main(){
	add := func (x,y int) int {
		return x+y
	}
		subtract := func (x, y int) int {
			return x-y
		
	}
	fmt.Println(apply(add, 2,3))
	fmt.Println(apply(subtract, 5,2))
}
func apply(f func(int, int) int, x,y int) int{
	return f(x,y)
}