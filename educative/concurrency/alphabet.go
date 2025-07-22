package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func lowerCase(goChannel chan<- string) {
	defer wg.Done()

	for ch := 'a'; ch <= 'z'; ch++ {
		goChannel <- fmt.Sprintf("%c", ch)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}

func upperCase(goChannel chan<- string) {
	defer wg.Done()

	for ch := 'A'; ch <= 'Z'; ch++ {
		goChannel <- fmt.Sprintf("%c", ch)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}

//var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(2)
	wg.Add(2)

	goChannel := make(chan string)
	go lowerCase(goChannel)
	go upperCase(goChannel)

	go func() {
		for val := range goChannel {
			fmt.Println(val)
		}
	}()

	wg.Wait()
}
