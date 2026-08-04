package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	files := []string{
		"report.pdf",
		"photo.jpg",
		"video.mp4",
		"archive.zip",
		"music.mp3",
		"document.docx",
		"image.png",
		"notes.txt",
	}

	workerCount := 3

	var wg sync.WaitGroup
	wg.Add(len(files))

	chSim := make(chan struct{}, workerCount)

	for _, file := range files {
		go func(filename string) {
			defer wg.Done()
			chSim <- struct{}{}

			defer func() {
				<- chSim
			}()

			ProcessFile(filename)
		}(file)
	}

	wg.Wait()
}

func ProcessFile(fileName string) {
	fmt.Printf("Началась обработка %s\n", fileName)
	time.Sleep(2 * time.Second)
	fmt.Printf("Закончилась обработка %s\n", fileName)
}
