// The goal of this activity is to explore the use of threads by creating a program for sorting integers that uses four goroutines to create four sub-arrays and then merge the arrays into a single array.
package main

import (
	"fmt"
	"sync"
)

// These Swap and BubbleSort is from previous assignment with modified
func Swap(array []int, i int) {
	temp := array[i]
	array[i] = array[i+1]
	array[i+1] = temp
}
func BubbleSort(array []int, wg *sync.WaitGroup) []int {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				Swap(array, j)
			}
		}
	}
	wg.Done()
	return array
}

func main() {
	var numCount int
	var numValue int
	var inputSlice = make([]int, 0)
	var waitGroup sync.WaitGroup

	fmt.Printf("Enter how many integer you want to insert: ")
	fmt.Scanf("%d", &numCount)

	for i := 0; i < numCount; i++ {
		fmt.Printf("Enter integer value #%d : ", i)
		fmt.Scanf("%d", &numValue)

		inputSlice = append(inputSlice, numValue)
	}

	// Partition the array into 4 parts
	size := numCount / 4
	part1 := inputSlice[:size]
	part2 := inputSlice[size : 2*(size)]
	part3 := inputSlice[2*(size) : 3*(size)]
	part4 := inputSlice[3*(size):]

	fmt.Println("Here are your arrays when partitioned", part1, part2, part3, part4)

	// Sort each array
	waitGroup.Add(4)
	go BubbleSort(part1, &waitGroup)
	go BubbleSort(part2, &waitGroup)
	go BubbleSort(part3, &waitGroup)
	go BubbleSort(part4, &waitGroup)
	waitGroup.Wait()

	// Final sort
	var finalSort = make([]int, 0)
	finalSort = append(finalSort, part1...)
	finalSort = append(finalSort, part2...)
	finalSort = append(finalSort, part3...)
	finalSort = append(finalSort, part4...)
	waitGroup.Add(1)
	go BubbleSort(finalSort, &waitGroup)
	waitGroup.Wait()

	fmt.Println(finalSort)
}
