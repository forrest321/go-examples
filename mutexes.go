package main

import (
	"fmt"
	"strconv"
	"sync"
)

// Basic demo of mutex
type mutexThing struct {
	sync.Mutex
	unsafeMap map[int]string
	safeMap   sync.Map
}

func mutexes() {
	m := &mutexThing{
		unsafeMap: make(map[int]string),
		safeMap:   sync.Map{},
	}
	m.unsafeMap[0] = "value 0"
	m.unsafeMap[1] = "value 1"
	m.safeMap.Store(0, "value 0")
	m.safeMap.Store(1, "value 1")
	fmt.Println("created mutexThing")
	fmt.Printf("unsafe map before: %+v\n", m.unsafeMap)
	fmt.Println("safe map before:")
	m.safeMap.Range(func(key, value interface{}) bool {
		fmt.Printf("key: %v, value: %v\n", key, value)
		return true
	})
	wg := sync.WaitGroup{}
	wg.Add(2)
	go doUnsafeWork(m, &wg)
	go doSafeWork(m, &wg)
	wg.Wait()
	fmt.Printf("unsafe map after: %+v\n", m.unsafeMap)
	fmt.Println("safe map after:")
	m.safeMap.Range(func(key, value interface{}) bool {
		fmt.Printf("key: %v, value: %v\n", key, value)
		return true
	})
	fmt.Println()
}

func doSafeWork(m *mutexThing, wg *sync.WaitGroup) {
	defer wg.Done()
	m.Lock()
	defer m.Unlock()

	m.unsafeMap[1] = "value " + strconv.Itoa(1)
	m.safeMap.Store(1, "value "+strconv.Itoa(1))

	m.unsafeMap[0] = "safely changed values"
	m.safeMap.Store(0, "changed values")
}

func doUnsafeWork(m *mutexThing, wg *sync.WaitGroup) {
	defer wg.Done()
	//for i := 1; i < 11; i++ {
	//	m.unsafeMap[i] = "don't do this"
	//}
	m.safeMap.Store(1, "this is fine")
}
