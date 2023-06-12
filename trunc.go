// The goal of this assignment is to practice working with floating point numbers and truncation code in Go.
package main

import (
	"fmt"
)

func main() {
	var number float64;
	fmt.Printf("Enter: ")
    _, err := fmt.Scanf("%f", &number)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Output: %d \n", int(number))
}
