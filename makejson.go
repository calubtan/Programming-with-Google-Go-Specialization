// The goal of this program is to practice working with RFCs in Go, using JSON as an example.
package main

import (
	"fmt"
	"encoding/json"
)

type Person struct {
	Name string
	Address string
}

func main () {
	var p1 Person;
	fmt.Printf("Enter name: ");
	_, e1 := fmt.Scanf("%s", &p1.Name);
	fmt.Printf("Enter address: ");
	_, e2 := fmt.Scanf("%s", &p1.Address);

	if e1 == nil && e2 == nil {
		earth, e3 := json.Marshal(p1);
		if e3 == nil {
			fmt.Println(string(earth));
		}
	}
}
