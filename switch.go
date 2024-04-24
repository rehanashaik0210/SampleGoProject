package main

import (
	"fmt"
)

// switch stmt is similar ti if else, But will stop once the condition satisfied.
// it will not carry forward to next statements after valid pass
func main() {
	switch {
	case false:
		fmt.Println("This should not print")

	case (2 == 4):
		fmt.Println("This should not print")
	case (3 == 3):
		fmt.Println("it should print")
	case (4 == 4):
		fmt.Println("Also true, does it print ? ")
	}

}
