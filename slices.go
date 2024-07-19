package main

import (
	"fmt"
	"math/rand"
	slices2 "slices"
	"time"
)

type signedInts interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

var sortTypes = []string{"bubble", "quick", "cocktail", "gnome", "counting", "heap", "default"}

// small int slice to show basic handling
// much of this is now included in std library
var ints []int = []int{1, 2, 3, 4, 5}

func slices() {
	fmt.Printf("ints slice: %v\n", ints)
	fmt.Printf("len is: %v\n", len(ints))
	fmt.Printf("cap is: %v\n", cap(ints))
	fmt.Printf("sorted: %v\n", isSorted(ints))
	shuffleAndSort(ints)
	bigSlice := makeBigSlice(15000)
	fmt.Printf("bigSlice len is: %v\n", len(bigSlice))
	fmt.Printf("bigSlice cap is: %v\n", cap(bigSlice))
	fmt.Printf("new bigSlice sorted: %v\n", isSorted(bigSlice))
	shuffleAndSort(bigSlice)
}

func shuffleAndSort[T signedInts](arr []T) {
	fmt.Println("Shuffling and sorting...")
	var startTime, endTime time.Time
	var d time.Duration
	times := make(map[string]time.Duration, len(arr))
	for _, st := range sortTypes {
		arr = shuffle(arr)
		startTime = time.Now()
		switch st {
		case "bubble":
			bubbleSort(arr)
		case "quick":
			quickSort(arr)
		case "cocktail":
			cocktailSort(arr)
		case "gnome":
			gnomeSort(arr)
		case "counting":
			countingSort(arr)
		case "heap":
			heapSort(arr)
		default:
			slices2.Sort(arr)
		}
		endTime = time.Now()
		d = endTime.Sub(startTime)
		times[st] = d
	}
	fmt.Println("Sorting complete")
	for k, v := range times {
		if len(arr) < 100 {
			fmt.Printf("%s sort took %v nanoseconds\n", k, v.Nanoseconds())
		} else {
			fmt.Printf("%s sort took %v microseconds\n", k, v.Microseconds())
		}
	}
}

func makeBigSlice(n int) []int32 {
	s := make([]int32, n)
	for i := range s {
		s[i] = rand.Int31()
	}
	return s
}

func shuffle[T signedInts](arr []T) []T {
	for i := len(arr) - 1; i > 0; i-- {
		if j := rand.Intn(i + 1); i != j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	return arr
}

func isSorted[T signedInts](arr []T) bool {
	//This is replaced by the slices std package
	//here for illustration purposes
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}

// Sorting functions found here:
// https://github.com/0xAX/go-algorithms/tree/master/sorting
// then converted to generics
func bubbleSort[T signedInts](arr []T) []T {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func quickSort[T signedInts](arr []T) []T {

	if len(arr) <= 1 {
		return arr
	}

	median := arr[rand.Intn(len(arr))]

	lowPart := make([]T, 0, len(arr))
	highPart := make([]T, 0, len(arr))
	middlePart := make([]T, 0, len(arr))

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

func cocktailSort[T signedInts](arr []T) []T {
	tmp := T(0)

	for i := 0; i < len(arr)/2; i++ {
		left := 0
		right := len(arr) - 1

		for left <= right {

			if arr[left] > arr[left+1] {
				tmp = arr[left]
				arr[left] = arr[left+1]
				arr[left+1] = tmp
			}

			left++

			if arr[right-1] > arr[right] {
				tmp = arr[right-1]
				arr[right-1] = arr[right]
				arr[right] = tmp
			}

			right--
		}
	}
	return arr
}

func gnomeSort[T signedInts](arr []T) []T {
	i := 1
	tmp := T(0)
	for i < len(arr) {
		if arr[i] >= arr[i-1] {
			i++
		} else {
			tmp = arr[i]
			arr[i] = arr[i-1]
			arr[i-1] = tmp

			if i > 1 {
				i--
			}
		}
	}
	return arr
}

func countingSort[T signedInts](arr []T) []T {
	k := getK(arr)
	arrayOfCounts := make([]T, k)

	for i := 0; i < len(arr); i++ {
		arrayOfCounts[arr[i]] += 1
	}

	for i, j := T(0), T(0); i < k; i++ {
		for {
			if arrayOfCounts[i] > 0 {
				arr[j] = i
				j += 1
				arrayOfCounts[i] -= 1
				continue
			}
			break
		}
	}
	return arr
}

func heapSort[T signedInts](arr []T) []T {
	i := 0
	tmp := T(0)

	for i = len(arr)/2 - 1; i >= 0; i-- {
		arr = sift(arr, T(i), T(len(arr)))
	}

	for i = len(arr) - 1; i >= 1; i-- {
		tmp = arr[0]
		arr[0] = arr[i]
		arr[i] = tmp
		arr = sift(arr, 0, T(i))
	}
	return arr
}

func getK[T signedInts](arr []T) T {
	if len(arr) == 0 {
		return 1
	}

	k := arr[0]

	for _, v := range arr {
		if v > k {
			k = v
		}
	}

	return k + 1
}

func sift[T signedInts](arr []T, i T, arrLen T) []T {
	done := false

	tmp := T(0)
	maxChild := T(0)

	for (i*2+1 < arrLen) && (!done) {
		if i*2+1 == arrLen-1 {
			maxChild = i*2 + 1
		} else if arr[i*2+1] > arr[i*2+2] {
			maxChild = i*2 + 1
		} else {
			maxChild = i*2 + 2
		}

		if arr[i] < arr[maxChild] {
			tmp = arr[i]
			arr[i] = arr[maxChild]
			arr[maxChild] = tmp
			i = maxChild
		} else {
			done = true
		}
	}

	return arr
}
