// The goals of this activity are to explore the design of concurrent algorithms and resulting synchronization issues.
package main

import (
	"fmt"
	"sync"
	"math/rand"
)

var wg sync.WaitGroup

type ChopS struct{ sync.Mutex }

type Philo struct {
	chopstick	[]*ChopS
	id			int
}

func (p Philo) eat(channel chan int) {
	// Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
   	for i:=0; i < 3; i++ {
		// Get permission
		channel <- 1

		// The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
		chopsL := rand.Intn(2)
        chopsR := (chopsL + 1) % 2

		// Lock
        p.chopstick[chopsL].Lock()
        p.chopstick[chopsR].Lock()

      	fmt.Println("starting to eat", p.id)

		// Unlock
        p.chopstick[chopsL].Unlock()
        p.chopstick[chopsR].Unlock()

	  	fmt.Println("finishing eating", p.id)
		
		// Release
		<- channel
   	}
	wg.Done()
}

func main() {
	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
   		CSticks[i] = new(ChopS)
	}

	// Each philosopher is numbered, 1 through 5.
	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
   		philos[i] = &Philo{
			   			[]*ChopS{CSticks[i], CSticks[(i+1)%5]},
			   			i + 1 }
	}

	// The host allows no more than 2 philosophers to eat concurrently.
	wg.Add(2)
	var channel = make(chan int, 2)
	for j := range philos {
		go philos[j].eat(channel)
	}

	wg.Wait()
}
