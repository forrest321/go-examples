package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"time"
)

// What should we count to? Try a million!
var countTo = 15000

// Keep channels in an object to easily pass as params
type Chans struct {
	countCh     chan int
	doneCh      chan struct{}
	interruptCh chan struct{}
	printCh     chan string
}

func concurrency() {
	// Get the number of currently available CPUs
	numGoroutines := runtime.GOMAXPROCS(0)

	// Useful for higher numbers. If you get tired of waiting
	// to count to 1 billion, ctrl-c will end it
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)

	// Lets see how long it takes
	started := time.Now()

	// Create a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	// Channel to collect current counts from each goroutine
	countCh := make(chan int, numGoroutines)
	// Channel to signal completion of counting
	doneCh := make(chan struct{})
	// Channel to signal interrupt
	interruptCh := make(chan struct{})
	// Channel to receive current string
	printCh := make(chan string)

	channels := Chans{countCh, doneCh, interruptCh, printCh}

	// Launch goroutines
	rangeSize := countTo / numGoroutines
	for i := 0; i < numGoroutines; i++ {
		start := i*rangeSize + 1
		end := (i + 1) * rangeSize
		// Ensure the last goroutine covers the full range
		if i == numGoroutines-1 {
			end = countTo
		}
		go count(start, end, i, &wg, channels)
	}

	// Goroutine to wait for all counts to complete
	go func() {
		wg.Wait()
		close(channels.doneCh)
	}()

	// Wait for signal or completion
	select {
	case <-sigs:
		close(channels.interruptCh)
		fmt.Println("Received interrupt signal")
	case msg := <-channels.printCh:
		fmt.Println(msg)
	case <-channels.doneCh:
		ended := time.Now()
		duration := ended.Sub(started)
		fmt.Printf("\nStarted: %v \nEnded: %v \nDuration: %v\n", started, ended, duration)
	}

	// Collect final counts
	highestCount := 0
	for i := 0; i < numGoroutines; i++ {
		c := <-countCh
		if c > highestCount {
			highestCount = c
		}
	}
	fmt.Printf("Used %v processors\n", numGoroutines)
	fmt.Printf("Counted to: %v\n", highestCount)
}

func count(start, end, chNum int, wg *sync.WaitGroup, chans Chans) {
	defer wg.Done()
	startTime := time.Now()
	currentNumber := 0
	for i := start; i <= end; i++ {
		select {
		case <-chans.interruptCh:
			chans.countCh <- currentNumber
			return
		default:
			currentNumber = i
			time.Sleep(1 * time.Microsecond)
		}
	}
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	fmt.Printf("Channel #%v counted from %v to %v in %v seconds\n", chNum, start, end, duration.Seconds())
	chans.countCh <- currentNumber

}
