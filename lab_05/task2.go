package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counterAtomic int64

func RunAtomicTask() {

	atomic.StoreInt64(&counterAtomic, 0)

	evenChan := make(chan int)
	oddChan := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()

		for num := range evenChan {
			if num%3 == 0 {
				atomic.AddInt64(&counterAtomic, 1)
			}
		}
	}()

	go func() {
		defer wg.Done()

		for num := range oddChan {
			if num%33 == 0 {
				atomic.AddInt64(&counterAtomic, -1)
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

	fmt.Println("Atomic counter:", counterAtomic)
}
