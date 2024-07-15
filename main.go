package main

import "fmt"

const separator string = "-=-=-=-=-=-=-=-=-=-=-=-"

func main() {
	fmt.Println("Hello World")
	separate("Basics")
	basics()
	separate("Slices")
	slices()
	separate("Inheritance")
	inheritance()
	fmt.Println("Mutexes")
	mutexes()
	separate("Concurrency")
	concurrency()
}

func separate(s string) {
	fmt.Println(separator)
	if len(s) > 0 {
		fmt.Println(s)
		fmt.Println(separator)
	}
}
