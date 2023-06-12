// The goal of this assignment is to practice working with integers, slices and loops in Go.
package main

import (
	"fmt"
	"strconv"
	"strings"
	"sort"
)

func main() {
	var input string;
  	s := make([]int, 0, 3)

	// Read first data
	fmt.Printf("Enter input integer: ");
	_, e1 := fmt.Scanf("%s", &input);
	if e1 == nil {
		for strings.ToUpper(input) != "X" {
			// Append data to slice
			in_data, e2 := strconv.Atoi(input);
			if e2 == nil {
				s = append(s, in_data);
			}

			// Print sorted slice
			sort.Sort(sort.IntSlice(s));
			fmt.Println("Sorted slice: ", s)

			// Read next data
			fmt.Printf("Enter input integer: ");
			_, e1 := fmt.Scanf("%s", &input);
			if e1 != nil {
				break;
			}
		}
  	}
}
