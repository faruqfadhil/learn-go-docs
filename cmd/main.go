package main

import (
	"fmt"
	"time"

	"github.com/faruqfadhil/learn-go-docs/channel"
)

func main() {
	workerPoolMain()
}

func bufferedMain() {
	ch := make(chan int, 2)
	go channel.Write(ch)
	time.Sleep(2 * time.Second)
	for v := range ch {
		fmt.Println("read value", v, "from ch")
		time.Sleep(2 * time.Second)
	}
}

func workerPoolMain() {
	startTime := time.Now()
	noOfJobs := 100
	go channel.Allocate(noOfJobs)

	done := make(chan bool)
	go channel.PrintResult(done)
	noOfWorkers := 10
	channel.CreateWorkerPool(noOfWorkers)

	<-done

	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}
