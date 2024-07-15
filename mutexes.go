package main

import (
	"fmt"
	"strconv"
	"sync"
)

//mutex
//sync
//go routines

type mutexThing struct {
	mu        sync.Mutex
	unsafeMap map[int]string
	safeMap   sync.Map
}

/*
Sync program
Goroutines and go channels
*/
func mutexes() {
	m := &mutexThing{
		unsafeMap: make(map[int]string),
		safeMap:   sync.Map{},
	}
	m.unsafeMap[0] = "value 0"
	m.unsafeMap[1] = "value 1"
	m.safeMap.Store(0, "value 0")
	m.safeMap.Store(1, "value 1")
	fmt.Printf("created mutexThing: %+v\n", m)

	go doUnsafeWork(m)
	go doSafeWork(m)

	fmt.Println()
}

func doSafeWork(m *mutexThing) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for i := 1; i < 11; i++ {
		m.unsafeMap[i] = "value " + strconv.Itoa(i)
		m.safeMap.Store(i, "value "+strconv.Itoa(i))
	}
	m.unsafeMap[0] = "safely changed values"
	m.safeMap.Store(0, "changed values")
}

func doUnsafeWork(m *mutexThing) {
	m.unsafeMap[0] = "don't do this"
	m.safeMap.Store(1, "this is fine")
}
