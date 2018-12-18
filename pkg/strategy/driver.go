package strategy

import "fmt"

// Run - Runs the package
func Run() {
	add := Operation{Addition{}}
	fmt.Println(add.Operate(3, 5))

	mul := Operation{Multiplication{}}
	fmt.Println(mul.Operate(3, 5))

}
