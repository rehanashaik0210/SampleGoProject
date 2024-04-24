package main
import "fmt"
func main(){
avg := [5] int {1,2,3,4,5.}
fmt.Println(avg)
sum :=0
for i:=0; i<5; i++{
	sum += avg [i]
}
arrayAvg :=sum/5
fmt.Println("The average of array items", arrayAvg)
fmt.Println("The sum of array items", sum)
}

