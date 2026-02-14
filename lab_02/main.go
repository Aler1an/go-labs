package main

import (
	"fmt"
	"reflect"
)

func main() {

	TaskSliceArray()

	shapes := []Shape{
		Circle{Radius: 5},
		Rectangle{Width: 4, Height: 6},
		Triangle{A: 3, B: 5, C: 4},
	}

	for _, shape := range shapes {
		typeName := reflect.TypeOf(shape).Name()
		fmt.Printf("%s: Area = %.2f, Perimeter = %.2f\n", typeName, shape.Area(), shape.Perimeter())
	}

}
