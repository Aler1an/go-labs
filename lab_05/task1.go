package main

import (
	"fmt"
	"sync"
)

var counter int
var mu sync.Mutex

func RunMutexTask() {

	counter = 0

	evenChan := make(chan int)
	oddChan := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	// обробка парних
	go func() {
		defer wg.Done()

		for num := range evenChan {
			if num%3 == 0 {
				mu.Lock()
				counter++
				mu.Unlock()
			}
		}
	}()

	// обробка непарних
	go func() {
		defer wg.Done()

		for num := range oddChan {
			if num%33 == 0 {
				mu.Lock()
				counter--
				mu.Unlock()
			}
		}
	}()

	for i := 1; i <= 1000; i++ {

		if i%2 == 0 {
			evenChan <- i
		} else {
			oddChan <- i
		}

	}

	close(evenChan)
	close(oddChan)

	wg.Wait()

	fmt.Println("Mutex counter:", counter)
}
