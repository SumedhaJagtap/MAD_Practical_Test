package main

import (
	"fmt"
	"sync"
	"time"
)

var mux = sync.Mutex{}

func allocateCust(custNo int, cashierNo int, timePerCustomer int, cashier chan<- int) {
	mux.Lock()
	cashier <- custNo
	fmt.Printf("Cashier %d: Customer %d Started", cashierNo, custNo)
	time.Sleep(3 * time.Second)
	mux.Unlock()
}

func main() {
	var wg = sync.WaitGroup{}
	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)
	timePerCustomer := 3
	noCust := 10
	wg.Add(3)
	for i := 1; i <= noCust; i++ {
		go func(i int) {
			defer wg.Done()
			if _, ok := <-c1; !ok {

				allocateCust(i, 1, timePerCustomer, c1)

			} else if _, ok := <-c2; !ok {

				allocateCust(i, 2, timePerCustomer, c1)

			} else if _, ok := <-c3; !ok {

				allocateCust(i, 3, timePerCustomer, c1)

			}
		}(i)
	}

	for i := 1; i <= 100; i++ {
		select {
		case msg1 := <-c1:
			fmt.Printf("Cashier 1: Customer %d Completed", msg1)
		case msg2 := <-c2:
			fmt.Printf("Cashier 2: Customer %d Completed", msg2)
		case msg3 := <-c3:
			fmt.Printf("Cashier 3: Customer %d Completed", msg3)
		}
	}
	wg.Wait()
}
