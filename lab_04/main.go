package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int, 10)
	ch3 := make(chan int)
	done := make(chan int)

	go generate(ch1)
	go filterEven(ch1, ch2)
	go square(ch2, ch3)
	go sum(ch3, done)

	total := <-done
	fmt.Println("Sum: ", total)
}

func generate(out chan int) {
	for i := 1; i <= 100; i++ {
		out <- i
	}

	close(out)
}

func filterEven(in chan int, out chan int) {
	for num := range in {
		if num%2 == 0 {
			out <- num
		}
	}

	close(out)
}

func square(in chan int, out chan int) {
	for num := range in {
		out <- num * num
	}

	close(out)
}

func sum(in chan int, done chan int) {
	total := 0
	for num := range in {
		total += num
	}
	done <- total
}
