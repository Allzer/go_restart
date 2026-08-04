package main

import (
	"fmt"
)

func main() {

	ch1 := make(chan string)
	ch2 := make(chan string)
	out := make(chan string)

	go GeneratorA(ch1)
	go GeneratorB(ch2)
	go Merge(ch1, ch2, out)

	for v := range out {
		fmt.Println(v)
	}
}

func GeneratorA(ch chan<- string) {
	for i := range 3 {
		ch <- fmt.Sprintf("A%d", i)
	}
	close(ch)
}

func GeneratorB(ch chan<- string) {
	for i := range 3 {
		ch <- fmt.Sprintf("B%d", i)
	}
	close(ch)
}

func Merge(ch1, ch2 <-chan string, out chan<- string) {
	ch1Closed := false
	ch2Closed := false
	for {
		select {
		case v, ok := <-ch1:
			if ok {
				out <- v
			} else {
				ch1Closed = true
				ch1 = nil
			}
		case v, ok := <-ch2:
			if ok {
				out <- v
			} else {
				ch2Closed = true
				ch2 = nil
			}
		}
		if ch1Closed && ch2Closed {
			close(out)
			break
		}
	}
}
