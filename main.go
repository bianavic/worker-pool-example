package main

import (
	"fmt"
	"time"
)

// stub
func worker(id int, jobs chan int, results chan int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j, "...")
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2 // send information to a result channel
	}
}

func main() {
	// 5 jobs to perform
	const numJobs = 5

	// channel for jobs
	jobs := make(chan int, numJobs)    // channel allows us to send a unit of work to worker pool
	results := make(chan int, numJobs) // define a channel named results that we're going to send results to channel, so each worker will send its results to this channel

	// 3 workers to consume jobs | here we spam 3 workers
	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)
	}

	// send jobs to channel
	for j := 1; j <= numJobs; j++ {
		jobs <- j // to our jobs channel send the value of J
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}

}
