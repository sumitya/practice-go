package main

import "fmt"

type Animal struct {
	Name   string
	Weight float32
}

func (a Animal) Eat() {
	fmt.Printf("%s is eating\n", a.Name)
}

func (a Animal) Sleep() {
	fmt.Printf("%s is sleeping\n", a.Name)
}

type Dog struct {
	Animal
	Breed string
}

func (d Dog) Bark() {
	fmt.Printf("%s is barking\n", d.Name)
}

func main() {
	// composition, Dog is an Animal
	bullDog := Dog{
		Animal: Animal{
			Name:   "Buddy",
			Weight: 16.5,
		},
		Breed: "BullDog",
	}

	fmt.Printf("Dog's Name is %s\n", bullDog.Name)
	fmt.Printf("Dog's Weight is %v\n", bullDog.Weight)

	bullDog.Eat()
	bullDog.Sleep()
	bullDog.Bark()
}
