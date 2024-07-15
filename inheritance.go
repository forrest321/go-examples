package main

import "fmt"

type animal struct {
	name     string
	legs     int
	fur      bool
	feathers bool
	sound    string
}

type dog struct {
	animal
	leash string
}

type cat struct {
	animal
	leash     string
	hairballs bool
}

type bird struct {
	animal
	wings int
}

type birdOfPrey struct {
	bird
	talons string
}

func inheritance() {
	generic := animal{name: "animal", legs: 0, fur: false, feathers: false}
	fmt.Printf("basic animal: %+v\n", generic)
	fido := dog{
		animal: animal{
			name:     "dog",
			legs:     4,
			fur:      true,
			feathers: false,
			sound:    "bark",
		},
		leash: "yaaay lets go outside",
	}
	fmt.Printf("fido: %+v\n", fido)
	felix := cat{
		animal: animal{
			name:     "cat",
			legs:     4,
			fur:      true,
			feathers: false,
			sound:    "meow",
		},
		leash:     "you will pay for this",
		hairballs: true,
	}
	fmt.Printf("felix: %+v\n", felix)
	eagle := birdOfPrey{
		bird: bird{
			animal: animal{
				name:     "eagle",
				legs:     2,
				fur:      false,
				feathers: true,
				sound:    "screech",
			},
			wings: 2,
		},
		talons: "sharp",
	}
	fmt.Printf("eagle: %+v\n", eagle)
}
