package calc

import (
	"errors"
	"fmt"
)

func init() {
	fmt.Println("Пакет ініціалізовано")
}

func Sum(nums ...float64) float64 {
	var total float64
	for _, num := range nums {
		total += num
	}

	return total
}

func Max(nums ...float64) float64 {

	if len(nums) == 0 {
		return 0
	}

	max := nums[0]
	for _, num := range nums[1:] {
		if num > max {
			max = num
		}
	}

	return max
}

func Min(nums ...float64) float64 {

	if len(nums) == 0 {
		return 0
	}

	min := nums[0]
	for _, num := range nums[1:] {
		if num < min {
			min = num
		}
	}

	return min
}

func Divide(a, b float64) (float64, error) {

	if b == 0 {
		return 0, errors.New("На нуль ділити не можна")
	}

	return a / b, nil
}

type Calculator interface {
	Sum(nums ...float64) float64
	Max(nums ...float64) float64
	Min(nums ...float64) float64
	Divide(a, b float64) (float64, error)
}

type Calc struct{}

func (c Calc) Sum(nums ...float64) float64 {
	return Sum(nums...)
}

func (c Calc) Max(nums ...float64) float64 {
	return Max(nums...)
}

func (c Calc) Min(nums ...float64) float64 {
	return Min(nums...)
}

func (c Calc) Divide(a, b float64) (float64, error) {
	return Divide(a, b)
}
