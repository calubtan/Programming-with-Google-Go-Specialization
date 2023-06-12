// The goal of this assignment is to create a routine that sorts a series of ten integers using the bubble sort algorithm.
package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func Swap(array []int, i int) {
	temp := array[i];
	array[i] = array[i+1];
	array[i+1] = temp;
}

func BubbleSort(array []int)[]int {
   	for i:=0; i< len(array)-1; i++ {
      	for j:=0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				Swap(array, j);
			}
      	}
   	}
   return array
}
func main() {
	fmt.Println("Please Enter list of integer (total up to 10): ")
	input := bufio.NewScanner(os.Stdin)
    input.Scan()
    input_string := input.Text()
	var numlist = strings.Split(input_string, " ")
	if len(numlist) > 10 {
		os.Exit(1);
	}

	// Filter the non-integer value
	var inSlice = make([]int, 0)
	for i := range numlist {
		in_integer, e2 :=strconv.Atoi(numlist[i]);
		if e2 == nil {
			inSlice = append(inSlice, in_integer);
		}
	}
   	fmt.Println(BubbleSort(inSlice))
}
