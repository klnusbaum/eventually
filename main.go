package main

import (
	"fmt"
	"sync"

	"github.com/gofrs/uuid"
	"github.com/klnusbaum/eventually/counters"
)

func main() {
	// we use go rountines to represent seperate processes/nodes
	// and channels to represent a network connection.
	// It doesn't seem worth it at the moment to actually
	// create seperate processes and network connections
	// as this is just a demo and is functionally equivalent.

	c12 := make(chan counters.GCounter, 1)
	c13 := make(chan counters.GCounter, 1)
	c21 := make(chan counters.GCounter, 1)
	c23 := make(chan counters.GCounter, 1)
	c31 := make(chan counters.GCounter, 1)
	c32 := make(chan counters.GCounter, 1)

	var allDone sync.WaitGroup
	allDone.Add(3)

	go doCount(
		&allDone,
		[]chan counters.GCounter{c12, c13},
		[]chan counters.GCounter{c21, c31},
	)
	go doCount(
		&allDone,
		[]chan counters.GCounter{c21, c23},
		[]chan counters.GCounter{c12, c32},
	)
	go doCount(
		&allDone,
		[]chan counters.GCounter{c31, c32},
		[]chan counters.GCounter{c13, c23},
	)

	allDone.Wait()
}

func doCount(allDone *sync.WaitGroup, txs, rxs []chan counters.GCounter) {
	defer allDone.Done()
	id := uuid.Must(uuid.NewV4())
	counter := counters.NewGCounter(id)
	counter.Inc()
	counter.Inc()
	counter.Inc()

	for _, tx := range txs {
		tx <- *counter
	}

	for _, rx := range rxs {
		received := <-rx
		counter.Merge(&received)
	}

	fmt.Printf("Hello from %s: %d\n", id, counter.Val())
}
