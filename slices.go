package main

import (
	"fmt"
	"math/rand"
	"time"
)

// small int slice to show basic handling
// much of this is now included in std library
var ints []int = []int{1, 2, 3, 4, 5}

func slices() {
	fmt.Printf("ints slice: %v\n", ints)
	fmt.Printf("len is: %v\n", len(ints))
	fmt.Printf("cap is: %v\n", cap(ints))
	fmt.Printf("sorted: %v\n", isSorted(ints))
	ints = shuffle(ints)
	fmt.Printf("shuffled slice: %v\n", ints)
	fmt.Printf("sorted: %v\n", isSorted(ints))
	ints = bubbleSort(ints)
	fmt.Printf("bubble sorted: %v\n", ints)
	fmt.Printf("sorted: %v\n", isSorted(ints))
	ints = shuffle(ints)
	fmt.Printf("shuffled slice: %v\n", ints)
	fmt.Printf("sorted: %v\n", isSorted(ints))
	ints = quickSort(ints)
	fmt.Printf("quick sorted: %v\n", ints)
	fmt.Printf("sorted: %v\n", isSorted(ints))
	bigSlice := makeBigSlice(15000)
	fmt.Printf("bigSlice len is: %v\n", len(bigSlice))
	fmt.Printf("bigSlice cap is: %v\n", cap(bigSlice))
	fmt.Printf("new bigSlice sorted: %v\n", isSorted32(bigSlice))
	startTime := time.Now()
	bigSlice = bubbleSort32(bigSlice)
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	fmt.Printf("bubble sort of bigSlice took: %v seconds\n", duration.Seconds())
	fmt.Printf("bigSlice sorted: %v\n", isSorted32(bigSlice))
	bigSlice = shuffle32(bigSlice)
	fmt.Printf("shuffled bigSlice sorted: %v\n", isSorted32(bigSlice))
	startTime2 := time.Now()
	bigSlice = quickSort32(bigSlice)
	endTime2 := time.Now()
	duration2 := endTime2.Sub(startTime2)
	fmt.Printf("quick sort of bigSlice took: %v seconds\n", duration2.Seconds())
	fmt.Printf("bigSlice sorted: %v\n", isSorted32(bigSlice))
	difference := duration - duration2
	fmt.Printf("quick sort faster by %v seconds\n", difference.Seconds())
}

func makeBigSlice(n int) []int32 {
	s := make([]int32, n)
	for i := range s {
		s[i] = rand.Int31()
	}
	return s
}

// The following methods can/should be genericized rather than
// creating separate methods based on int type
func shuffle(arr []int) []int {
	for i := len(arr) - 1; i > 0; i-- {
		if j := rand.Intn(i + 1); i != j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	return arr
}

func shuffle32(arr []int32) []int32 {
	for i := len(arr) - 1; i > 0; i-- {
		if j := rand.Intn(i + 1); i != j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	return arr
}

func isSorted(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}

func isSorted32(arr []int32) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}

// Sorting functions found here:
// https://github.com/0xAX/go-algorithms/tree/master/sorting
func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func bubbleSort32(arr []int32) []int32 {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func quickSort(arr []int) []int {

	if len(arr) <= 1 {
		return arr
	}

	median := arr[rand.Intn(len(arr))]

	lowPart := make([]int, 0, len(arr))
	highPart := make([]int, 0, len(arr))
	middlePart := make([]int, 0, len(arr))

	for _, item := range arr {
		switch {
		case item < median:
			lowPart = append(lowPart, item)
		case item == median:
			middlePart = append(middlePart, item)
		case item > median:
			highPart = append(highPart, item)
		}
	}

	lowPart = quickSort(lowPart)
	highPart = quickSort(highPart)

	lowPart = append(lowPart, middlePart...)
	lowPart = append(lowPart, highPart...)

	return lowPart
}

func quickSort32(arr []int32) []int32 {

	if len(arr) <= 1 {
		return arr
	}

	median := arr[rand.Intn(len(arr))]

	lowPart := make([]int32, 0, len(arr))
	highPart := make([]int32, 0, len(arr))
	middlePart := make([]int32, 0, len(arr))

	for _, item := range arr {
		switch {
		case item < median:
			lowPart = append(lowPart, item)
		case item == median:
			middlePart = append(middlePart, item)
		case item > median:
			highPart = append(highPart, item)
		}
	}

	lowPart = quickSort32(lowPart)
	highPart = quickSort32(highPart)

	lowPart = append(lowPart, middlePart...)
	lowPart = append(lowPart, highPart...)

	return lowPart
}
