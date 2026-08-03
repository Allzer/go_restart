package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	files := []string{
		"report.pdf",
		"photo.jpg",
		"video.mp4",
	}
	wg.Add(len(files))

	for _, v := range files {
		go func(file string) {
			defer wg.Done()
			ProcessFile(file)
		}(v)
	}
	wg.Wait()
}

func ProcessFile(filename string) {
	fmt.Printf("Работа с файлом %s началась\n", filename)
	time.Sleep(time.Second * 2)
	fmt.Printf("Работа с файлом %s закончилась\n", filename)
}
