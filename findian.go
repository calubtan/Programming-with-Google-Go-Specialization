// The goal of this assignment is to practice working with strings in Go.
package main

import (
	"fmt"
	"strings"
	"bufio"
  	"os"
)

func main(){
	input := bufio.NewScanner(os.Stdin)
    input.Scan()
    input_string := input.Text()
	lower_string := strings.ToLower(input_string);
	var string_length = len([]rune(lower_string));
	first_char := lower_string[0];
	last_char := lower_string[string_length-1];

	if  strings.Contains(string(first_char), "i") && 
		strings.Contains(lower_string, "a") &&
		strings.Contains(string(last_char), "n") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
