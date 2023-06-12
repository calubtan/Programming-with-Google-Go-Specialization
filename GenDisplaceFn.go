// The goal of this assignment is to create a routine that solves a linear kinematics problem.
package main

import (
	"fmt"
)

func GenDisplaceFn(a, v_o, s_o float64) func (float64) float64 {
	fn := func (t float64) float64 {
		return ( ((a * (t * t)) / 2) + (v_o * t) + s_o )
	}
	return fn
}

func main() {
	var a, v_o, s_o, t, s float64;
	fmt.Printf("Please enter acceleration (a): ")
	fmt.Scanf("%f", &a)
	fmt.Printf("Please enter initial velocity (v_o): ")
	fmt.Scanf("%f", &v_o)
	fmt.Printf("Please enter initial displacement (s_o): ")
	fmt.Scanf("%f", &s_o)

	fn := GenDisplaceFn(a, s_o, v_o)

	fmt.Printf("Please enter time (t): ")
	fmt.Scanf("%f", &t)

	s = fn(t)

	fmt.Printf("Displament after time %f is %f\n", t, s)
}
