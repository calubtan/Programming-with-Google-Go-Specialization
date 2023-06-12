// The goal of this activity is to explore race conditions by creating and running two simultaneous goroutines.
package main

var a int = 0
var b int = 0

// Race condition occured when both "access" and "write" function try to access
// read and write on variable a. It might not get reflected on latest value

func access() {
	for {
		var local_a int = a
		if local_a%2 == 0 {
			b = b + 1
		}
	}
}

func write(value int) {
	for {
		a = a + value
	}
}

func main() {
	go access()
	go write(1)
}
