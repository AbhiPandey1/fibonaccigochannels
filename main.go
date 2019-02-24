package main

import "fmt"

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	//Worker pools you can see the program utilize over 100% of cpu showing that milticore utilization
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)

	for index := 0; index < 100; index++ {
		jobs <- index
	}
	close(jobs)

	for result := range results {
		fmt.Println(result)
	}
}

func worker(jobs <-chan int, results chan<- int) {

	for job := range jobs {
		results <- fib(job)
	}
}
func fib(n int) int {
	if n <= 1 {
		return n
	}

	return fib(n-1) + fib(n-2)

}
