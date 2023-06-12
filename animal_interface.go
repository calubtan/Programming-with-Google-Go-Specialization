// The goal of this assignment is to write a GoLang routine that allows users to create a set of animals and then get information about those animals.
package main

import (
	"fmt"
	"strings"
)
type Animal interface {
	Eat() string
	Move() string
	Speak() string
}

type Cow struct{ food, locomotion, noise string }
func (c Cow) Eat() string {
	return c.food
}
func (c Cow) Move() string {
	return c.locomotion
}
func (c Cow) Speak() string{
	return c.noise
}

type Bird struct{ food, locomotion, noise string }
func (b Bird) Eat() string {
	return b.food
}
func (b Bird) Move() string {
	return b.locomotion
}
func (b Bird) Speak() string{
	return b.noise
}

type Snake struct{ food, locomotion, noise string }
func (s Snake) Eat() string {
	return s.food
}
func (s Snake) Move() string {
	return s.locomotion
}
func (s Snake) Speak() string{
	return s.noise
}

func createNewAnimal (input_1, input_2 string, animalMap map[string]Animal) {
	switch input_2 {
	case "cow":
		animal := Cow{food: "grass", locomotion: "walk", noise: "moo"}
		animalMap[input_1] = animal
		fmt.Println("Created it!")
		break
	case "bird":
		animal := Bird{food: "worms", locomotion: "fly", noise: "peep"}
		animalMap[input_1] = animal
		fmt.Println("Created it!")
		break
	case "snake":
		animal := Bird{food: "mice", locomotion: "slither", noise: "hsss"}
		animalMap[input_1] = animal
		fmt.Println("Created it!")
		break
	default:
		fmt.Println("Invalid animal")
		break
	}
}

func queryAnimal(input_1, input_2 string, animalMap map[string]Animal) {
	animal, err := animalMap[strings.ToLower(input_1)]
	if err {
		switch strings.ToLower(input_2) {
		case "eat":
			fmt.Println(animal.Eat())
			break
		case "move":
			fmt.Println(animal.Move())
			break
		case "speak":
			fmt.Println(animal.Speak())
			break
		default:
			fmt.Println("Invalid animal/action")
			break
		}
	} else {
		fmt.Println("Query name not in list");
	}
}

func main() {
	var input_1, input_2 string
	var action string
	var animalMap map[string]Animal
	animalMap = make(map[string]Animal)
	fmt.Println("> Usage:")
	fmt.Println(" 1. newanimal <animal name> <animal type (cow/bird/snake)>")
	fmt.Println(" 2. query <animal name> <animal action>")

	for {
		fmt.Printf("> ")
		fmt.Scanf("%s %s %s", &action, &input_1, &input_2)

		if strings.ToLower(action) == "newanimal" {
			createNewAnimal(input_1, input_2, animalMap)
		} else if strings.ToLower(action) == "query" {
			queryAnimal(input_1, input_2, animalMap)
		} else {
			fmt.Println("Invalid command")
		}

		//clean variable before next input
		action = ""
		input_1 = ""
		input_2 = ""
	}
}
