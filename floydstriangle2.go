package main
import "fmt"
func main(){
	var i, j, rows int
	var number int
fmt.Print("Enter the rows to print floyds triangle=")
fmt.Scanln(&rows)
fmt.Print("Enter the starting number=")
fmt.Scanln(&number)
fmt.Println("floyds triangle")
for i=1; i<=rows; i++{
	for j=1; j<=i; j++{
		fmt.Printf(" %d ", number)
		number++
	}
	fmt.Println()
}
}