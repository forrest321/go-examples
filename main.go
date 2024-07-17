// This is a demo of common Go usage. It is not meant to show
// how production code should be organized.
package main

import "fmt"

const separator string = "-=-=-=-=-=-=-=-=-=-=-=-"

func main() {
	//Why not start with a classic?
	fmt.Println("Hello World")
	separate("Basics")
	basics()
	separate("Slices")
	slices()
	separate("Inheritance")
	inheritance()
	separate("Interfaces")
	interfaces()
	separate("Mutexes")
	mutexes()
	separate("Concurrency")
	concurrency()
	separate("REST")
	rest()
}

func separate(s string) {
	//Adding simple borders to output
	fmt.Println(separator)
	if len(s) > 0 {
		fmt.Println(s)
		fmt.Println(separator)
	}
}
