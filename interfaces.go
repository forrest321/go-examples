package main

import "fmt"

// Similar to inheritance via structs, interfaces allow
// different structs to share functionality
type walkable interface {
	attachLeash()
	walk()
}
type talkable interface {
	talk()
}

func (d *dog) attachLeash() {
	fmt.Println(d.leash)
}
func (d *dog) walk() {
	fmt.Println("all goes well")
}
func (d *dog) talk() {
	fmt.Println(d.sound)
}
func (c *cat) attachLeash() {
	fmt.Println(c.leash)
}
func (c *cat) walk() {
	fmt.Println("this was a bad idea")
}
func (c *cat) talk() {
	fmt.Println(c.sound)
}
func (b *bird) talk() {
	fmt.Println(b.sound)
}

func interfaces() {
	speak(&fido)
	walkPet(&fido)
	speak(&felix)
	walkPet(&felix)
	//cant walk an eagle, won't compile
	//walkPet(&eagle)
	speak(&eagle)
}
func speak(pet talkable) {
	pet.talk()
}

func walkPet(pet walkable) {
	pet.attachLeash()
	pet.walk()
}
