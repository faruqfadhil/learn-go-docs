package channel

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Job represent the job.
type Job struct {
	ID       int
	Randomno int
}

// Result represent the result.
type Result struct {
	Job         Job
	SumOfDigits int
}

// Initial buffered channel.
var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

func digits(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	time.Sleep(2 * time.Second)
	return sum
}

// Worker is a worker func.
func Worker(wg *sync.WaitGroup) {
	for job := range jobs {
		output := Result{
			Job:         job,
			SumOfDigits: digits(job.Randomno),
		}
		results <- output
	}
	wg.Done()
}

// CreateWorkerPool is create of worker pool.
func CreateWorkerPool(noOfWorker int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorker; i++ {
		wg.Add(1)
		go Worker(&wg)
	}
	wg.Wait()
	close(results)
}

// Allocate allocate func.
func Allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randomNo := rand.Intn(999)
		job := Job{
			ID:       i,
			Randomno: randomNo,
		}
		jobs <- job
	}
	close(jobs)
}

// PrintResult print a result func.
func PrintResult(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.Job.ID, result.Job.Randomno, result.SumOfDigits)
	}
	done <- true
}
