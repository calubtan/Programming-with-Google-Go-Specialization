// The goal of this assignment is to practice working with the ioutil and os packages in Go.
package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

type Name struct {
    fname string
    lname string
}

func main () {
  	s := make([]Name, 0, 3);
	var filename string;
	var in1 Name;
	fmt.Printf("Please specify filename: ");
	_, e1 := fmt.Scanf("%s", &filename);
	if e1 == nil {
		f,e2 := os.Open(filename);
		if e2 == nil {
			fileread := bufio.NewScanner(f)
    		fileread.Split(bufio.ScanLines)
			for fileread.Scan() {
				first_last := strings.Fields(fileread.Text());
				in1.fname = first_last[0];
				in1.lname = first_last[1];
				s = append(s, in1);
    		}
		}
	}

	// Iterate thru slide to print out 
	for i := range s {
		fmt.Println(s[i]);
	}
}
