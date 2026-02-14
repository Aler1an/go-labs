package main

import (
	"fmt"
	"math/rand"
)

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
func main() {

	var y float64
	x := rand.Int31n(1000)

	if x > 100 {
		y = 0.5*float64(x) + float64(x)*float64(x)
	} else {
		if x != 0 {
			y = 12/float64(-x) + float64(x)*float64(x)
		} else {
			fmt.Println("На 0 ділити не можна")
			return
		}
	}

	fmt.Println("x: %d\n", x)
	fmt.Println("Result: %d\n", y)

	SecondVariant()
}
