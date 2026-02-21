package main

import (
	"fmt"
	"github.com/Aler1an/go-labs/lab_03/calc"
)

func main() {

	fmt.Println("Sum: ", calc.Sum(1, 2, 3, 4, 5))
	fmt.Println("Max: ", calc.Max(10, 3, 7, 25, 1))
	fmt.Println("Min: ", calc.Min(10, 3, 7, 25, 1))

	result, err := calc.Divide(10, 2)

	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Result: ", result)
	}

	_, err = calc.Divide(10, 0)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	var myCalc calc.Calc = calc.Calc{}

	fmt.Println("Sum: ", myCalc.Sum(1, 2, 3, 4, 5))
	fmt.Println("Max: ", myCalc.Max(10, 3, 7, 25, 1))
	fmt.Println("Min: ", myCalc.Min(10, 3, 7, 25, 1))

	res, err := myCalc.Divide(20, 4)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Result: ", res)
	}
}
