package main

import (
	"fmt"
	"sync"
	"time"
)

var mux = sync.Mutex{}

func allocateCust(custNo int, cashierNo int, timePerCustomer int, cashier chan<- int, cashier1 <-chan int) {
	mux.Lock()
	cashier <- custNo
	fmt.Printf("Cashier %d: Customer %d Started", cashierNo, custNo)
	no := <-cashier1
	time.Sleep(timePerCustomer * time.Second)
	fmt.Printf("Cashier %d: Customer %d Completed", cashierNo, custNo)
	mux.Unlock()
}

func main() {
	var wg = sync.WaitGroup{}
	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)
	timePerCustomer := 3

	wg.Add(3)
	go func(i int) {
		defer wg.Done()
		fmt.Printf("Cashier 1: Customer %d Started", i)
		c1 <- i
	}(i)

	go func(i int) {
		defer wg.Done()
		fmt.Printf("Cashier 2: Customer %d Started", i)
		c3 <- i
	}(i)
	go func(i int) {
		wg.Done()
		fmt.Printf("Cashier 3: Customer %d Started", i)
		c2 <- i
	}(i)

	wg.Close()
	for i := 1; i <= 100; i++ {
		defer wg.Done()
	}
}
