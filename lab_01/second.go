package main

import (
	"fmt"
	"math/rand"
)

func SecondVariant() {

	var y float64
	x := rand.Int31n(1000)

	switch {
	case x > 100:
		y = 0.5*float64(x) + float64(x)*float64(x)
	default:
		y = 12/float64(-x) + float64(x)*float64(x)
	}

	fmt.Println("x: %d\n ", x)
	fmt.Println("Result: %d\n ", y)
}
