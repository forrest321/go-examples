package main

import (
	"fmt"
	"math"
	"math/rand/v2"
)

var a string

const b string = "I'm a constant"

type c struct {
	x           int
	y           int
	name        string
	Exported    string
	notexported string
}

func basics() {
	a = "package level variable"
	fmt.Println(a)
	//b = "invalid assignment"
	fmt.Println(b)
	d := c{
		x:           1,
		y:           2,
		name:        "b",
		Exported:    "Exported",
		notexported: "Not Exported",
	}
	fmt.Println(d)

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

	if e > f {
		fmt.Println("e is greater than f")
	} else {
		fmt.Println("e is not greater than f")
	}

	switch e {
	case 1:
		fmt.Println("e == 1")
	case 2, 3:
		fmt.Println("e == 1 or e == 2")
	default:
		fmt.Println("e is something else")
	}

	fmt.Println("For loop")
	for i := 0; i < 10; i++ {
		if i == 0 {
			fmt.Printf(fmt.Sprintf("i=%v", i))
			continue
		}
		fmt.Printf(fmt.Sprintf(",%v", i))
	}
	fmt.Println()

	fmt.Println("Range")
	stringToRange := "It is easy to range over a string as a byte slice"
	for _, s := range stringToRange {
		fmt.Printf("%s", string(s))
	}
	fmt.Println()

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

func sqrt(x int) int {
	return int(math.Floor(float64(x) / math.Sqrt(float64(x))))
}
