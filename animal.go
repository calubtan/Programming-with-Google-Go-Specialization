// The goal of this assignment is to programmatically get information about a set of predefined objects using Go. Click the "My Submission" tab to view specific instructions for this activity.
package main

import "fmt"

type Animal struct {
	food string
	locomotion string
	noise string
}

func (p *Animal) InitMe(f, l, n string) {
	p.food = f
	p.locomotion = l
	p.noise = n
}
func (p *Animal) Eat() {
	fmt.Println(p.food)
}
func (p *Animal) Move() {
	fmt.Println(p.locomotion)
}
func (p *Animal) Speak() {
	fmt.Println(p.noise)
}

func PrintAction (p *Animal, action string) {
	switch action {
	case "eat":
		p.Eat()
	case "move":
		p.Move()
		break
	case "speak":
		p.Speak()
		break
	default:
		fmt.Println("Invalid action")
		break;
	}
}

func main() {
	cow := new(Animal)
	bird := new(Animal)
	snake := new(Animal)
	cow.InitMe("grass", "walk", "moo")
	bird.InitMe("worms", "fly", "peep")
	snake.InitMe("mice", "slither", "hsss")
	var animal_type string
	var action string
	for {
		fmt.Printf("> ")
		fmt.Scanf("%s %s", &animal_type, &action)
		// Check input
		switch animal_type {
		case "cow":
			PrintAction(cow, action)
			break
		case "bird":
			PrintAction(bird, action)
			break
		case "snake":
			PrintAction(snake, action)
			break
		default:
			fmt.Println("Invalid animal")
			break
		}
		//clean variable before next input
		animal_type = ""
		action = ""
	}
}
