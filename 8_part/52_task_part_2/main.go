package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	workersCount := 3
	var wg sync.WaitGroup
	wg.Add(workersCount)

	ctx, cancel := context.WithCancel(context.Background())

	for i := range workersCount {
		go Worker(i+1, &wg, ctx)
	}
	time.Sleep(5*time.Second)
	cancel()
	wg.Wait()
}

func Worker(id int, wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	ticker := time.NewTicker(time.Second * 1)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Работа Worker %d завершена\n", id)
			return
		case <-ticker.C:
			fmt.Printf("Worker %d работает\n", id)
		}
	}
}
