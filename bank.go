package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

func getCurrentTime() string {
	currentTime := time.Now()
	return currentTime.Format("2006-01-02 15:04:05")
}

var mux = sync.Mutex{}

func cashier(id int, jobs <-chan int, results chan<- int, timePerCustomer time.Duration) {
        /*
            This function will take channel id,jobs(sender channel),results(receiver channel), timePerCustomer as an input.
            It will iterate over all jobs of a each sending channel coming as parameter and will wait for timePerCustomer seconds to finish that job.

        */
        for j := range jobs {
		// mux.Lock()
		fmt.Println(getCurrentTime(), " --> Cashier ", id, ": Customer ", j, " Started")
		time.Sleep(time.Second * timePerCustomer)
		fmt.Println(getCurrentTime(), " --> Cashier ", id, ": Customer ", j, " Completed")
		results <- j * 2
		// fmt.Println("-----------------------------------------------------------------------")
		// mux.Unlock()
	}
}

func findValue(arg string) (string, string) {
        /*
            This function will take each command line argument as an input.
            Return key and value.
            Ex., arg: --numCashiers=3
                 return numCashiers,3

        */
	arg = strings.Replace(arg, "--", "", -1)
	r, _ := regexp.Compile("=[0-9]+")
	searched := r.FindString(arg)
	key := strings.Replace(arg, searched, "", -1)
	value := strings.Replace(searched, "=", "", -1)
	return key, value
}
func main() {
	variables := make(map[string]int)

	for _, i := range os.Args[1:] {
		k, v := findValue(i)
		variables[k], _ = strconv.Atoi(v)
	}

	chanJob := make(chan int, 100)      //channel to send customer
	results := make(chan int, 100)      //channel to invoke customer

	for w := 1; w <= variables["numCashiers"]; w++ {
		go cashier(w, chanJob, results, time.Duration(variables["timePerCustomer"]))
	}

	fmt.Println(getCurrentTime(), " --> Bank Simulation Started")

	for j := 1; j <= variables["numCustomers"]; j++ {
		chanJob <- j
	}

	close(chanJob)

	for j := 1; j <= variables["numCustomers"]; j++ {
		<-results
	}

	fmt.Println(getCurrentTime(), " --> Bank Simulated Ended")
}
