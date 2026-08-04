package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	jobs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	jobsChan := make(chan int)

	var wg sync.WaitGroup
	workersCount := 3
	wg.Add(workersCount)

	for i := 0; i < workersCount; i++ {
		go Worker(i, jobsChan, &wg)
	}

	for _, job := range jobs {
		jobsChan <- job
	}
	close(jobsChan)
	wg.Wait()
	fmt.Println("Все задачи выполнены")

}

func Worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d выполняет задачу %d\n", id, job)
		time.Sleep(time.Second * 2)
		fmt.Printf("Worker %d закончил задачу %d\n", id, job)
	}
}
