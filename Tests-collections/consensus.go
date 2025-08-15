package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func participant(id int, resultChan chan<- int, consensusChan <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	randNum := rand.Intn(100000)
	fmt.Printf("Participant %d picked %d\n", id, randNum)
	resultChan <- randNum // Send picked number to coordinator

	consensus := <-consensusChan // Wait for broadcasted consensus
	fmt.Printf("Participant %d received consensus: %d\n", id, consensus)
}

func findConsensus(results <-chan int, consensusChans []chan int) {
	minValue := int(^uint(0) >> 1) // max int

	// Collect all values and find the minimum
	for i := 0; i < len(consensusChans); i++ {
		num := <-results
		if num < minValue {
			minValue = num
		}
	}

	// Broadcast to all participants
	for _, ch := range consensusChans {
		ch <- minValue
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	const numParticipants = 10
	results := make(chan int, numParticipants)
	consensusChans := make([]chan int, numParticipants)
	var wg sync.WaitGroup

	// Create individual consensus channels for each participant
	for i := 0; i < numParticipants; i++ {
		consensusChans[i] = make(chan int, 1)
	}

	// Start participants
	for i := 0; i < numParticipants; i++ {
		wg.Add(1)
		go participant(i, results, consensusChans[i], &wg)
	}

	// Start coordinator to find consensus and broadcast
	go findConsensus(results, consensusChans)

	// Wait for all participants to complete
	wg.Wait()
}
