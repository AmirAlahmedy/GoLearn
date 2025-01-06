package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Printf("Processing job %d\n", j)
			} else {
				fmt.Println("All jobs processed")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
        jobs <- j
		fmt.Printf("Sent job %d\n", j)
    }
	close(jobs)
	fmt.Println("Sent all jobs")

	<-done

	_, ok := <-jobs
	fmt.Println("received more jobs: ", ok)

}