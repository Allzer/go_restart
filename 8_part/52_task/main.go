package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) //одна конструкция
	defer cancel()                                                          //одна конструкция

	var wg sync.WaitGroup
	wg.Add(1)

	go func(ctx context.Context) {
		defer wg.Done()
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Горутина остановлена")
				return
			case <-ticker.C:
				fmt.Println("Работаю...")
			}
		}
	}(ctx)

	wg.Wait()
}
