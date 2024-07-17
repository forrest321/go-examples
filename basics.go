package main

import (
	"fmt"
	"math"
	"math/rand/v2"
)

// Variables, constants, structs, etc., declared here
// will have scope throughout this package
var a string

const b string = "I'm a constant"

type c struct {
	x           int
	y           int
	name        string
	Exported    string //Case of the first letter determines visibility
	notexported string //lower case is not available outside package
}

func basics() {
	a = "package level variable"
	fmt.Println(a)
	//Reassigning a const is not allowed, so the following won't compile
	//b = "invalid assignment"
	fmt.Println(b)
	d := c{
		x:           1,
		y:           2,
		name:        "b",
		Exported:    "Exported",
		notexported: "Not Exported",
	}
	//There are multiple ways to print to console
	//In a production system, this would be logged instead
	fmt.Printf("%+v\n", d)

	//The compiler can determine the type
	//Also, init/assign of the same types can be on a single line
	e, f, g, h := 1, 2, 3, 4
	fmt.Printf("e, f, g, h := %v, %v, %v, %v\n", e, f, g, h)
	e, f, g, h = h, g, f, e
	fmt.Println("e, f, g, h = h, g, f, e")
	fmt.Println(e, f, g, h)
	e, f, g, h = e*e, f*f, g*g, h*h
	fmt.Println("e, f, g, h = e*e, f*f, g*g, h*h")
	fmt.Println(e, f, g, h)
	e, f, g, h = sqrt(e), sqrt(f), sqrt(g), sqrt(h)
	fmt.Println("e, f, g, h = e/√e, f/√f, g/√g, h/√h")
	fmt.Println(e, f, g, h)
	e, f, g, h = h, g, f, e
	fmt.Println("e, f, g, h = h, g, f, e")
	fmt.Println(e, f, g, h)

	//if statements are simple and common
	if e > f {
		fmt.Println("e is greater than f")
	} else {
		fmt.Println("e is not greater than f")
	}

	//switches are great for many things
	switch e {
	case 1:
		fmt.Println("e == 1")
	case 2, 3:
		fmt.Println("e == 1 or e == 2")
	default:
		fmt.Println("e is something else")
	}

	//For loops are useful for many things, like iterating over a slice
	fmt.Println("For loop")
	for i := 0; i < 10; i++ {
		if i == 0 {
			fmt.Printf(fmt.Sprintf("i=%v", i))
			continue
		}
		fmt.Printf(fmt.Sprintf(",%v", i))
	}
	fmt.Println()

	//Ranging works well for iterating over slices, or bytes in a string in this case
	fmt.Println("Range")
	stringToRange := "It is easy to range over a string as a byte slice"
	for _, s := range stringToRange {
		fmt.Printf("%s", string(s))
	}
	fmt.Println()

	//Technically a for loop, the "while" is great for less defined end cases
	fmt.Println("While loop")
	done := false
	match := 5
	for !done {
		check := rand.IntN(20)
		fmt.Printf("random number = %v\n", check)
		if check == match {
			fmt.Println("It matches!")
			done = true
		}
	}

	fmt.Println()
}

// helper functions reduce repetitive code
func sqrt(x int) int {
	return int(math.Floor(float64(x) / math.Sqrt(float64(x))))
}
