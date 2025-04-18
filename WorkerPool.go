package main

import (
	"fmt"
	"sync"
	"time"
	"runtime"
)

type SqError struct {
	Num    int
	Reason string
}

func (s SqError) Error() string {
	return fmt.Sprintf("Square Error with your number %d\nReason:%s", s.Num, s.Reason)
}

func Worker(jobs chan int, n int, result chan int, errorCh chan SqError, wg *sync.WaitGroup) {
	defer wg.Done()
	//fmt.Printf("The worker number %d started it's work\n", n)
	time.Sleep(500 * time.Millisecond)
	for val := range jobs {
		if val > 50 {
			errorCh <- SqError{
				Num:    val,
				Reason: "You cannot square more than 50!",
			}
		} else {
			//Squaring the values and putting it into the result once again
			result <- val * val
		}
	}
	fmt.Printf("The worker number %d Finished! \n", n)
}

func main() {
    runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	nJobs := 100
	nWorkers := 10
	jobs := make(chan int, nJobs)
	result := make(chan int, nJobs)
	errorCh := make(chan SqError, 50)

	//worker goroutines started and waiting for jobs to be fed
	for i := 1; i <= nWorkers; i++ {
		wg.Add(1)
		go Worker(jobs, i, result, errorCh, &wg)
	}

	//Jobs feeding into jobs channel
	for i := 1; i <= nJobs; i++ {
		jobs <- i
	}

	close(jobs)

	wg.Wait()
	close(result)
	close(errorCh)

	//Result printing
	for res := range result{
	    fmt.Println(res)
	}

	//error channel
	for err := range errorCh {
		fmt.Println("________Squaring Error________\n", err)
	}

	fmt.Println("Done!")
}