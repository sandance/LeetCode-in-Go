package main

import "fmt"

func main() {

	generator := func(done <-chan interface{}, integers ...int) <-chan int {
		initStream := make(chan int)
		go func() {
			defer close(initStream)
			for _, i := range integers {
				select {
				case <-done:
					return
				case initStream <- i:
				}
			}
		}()
		return initStream
	}

	multiply := func(done <-chan interface{}, initStream <-chan int, multiplier int) <-chan int {
		multiplyStream := make(chan int)
		go func() {
			defer close(multiplyStream)
			for i := range initStream {
				select {
				case <-done:
					return
				case multiplyStream <- i * multiplier:

				}
			}
		}()

		return multiplyStream
	}

	add := func(done <-chan interface{}, initStream <-chan int, additive int) <-chan int {
		addstream := make(chan int)
		go func() {
			defer close(addstream)
			for i := range initStream {
				select {
				case <-done:
					return
				case addstream <- i + additive:
				}
			}
		}()
		return addstream
	}

	done := make(chan interface{})
	defer close(done)

	initStream := generator(done, 1, 2, 3, 4)
	pipeline := multiply(done, add(done, multiply(done, initStream, 2), 1), 3)

	for v := range pipeline {
		fmt.Println(v)
	}

}
