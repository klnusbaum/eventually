package main

import (
	"fmt"
	"sync"

	"github.com/gofrs/uuid"
	"github.com/klnusbaum/eventually/counters"
)

func main() {
	c12 := make(chan counters.Payload, 1)
	c13 := make(chan counters.Payload, 1)
	c21 := make(chan counters.Payload, 1)
	c23 := make(chan counters.Payload, 1)
	c31 := make(chan counters.Payload, 1)
	c32 := make(chan counters.Payload, 1)

	var allDone sync.WaitGroup
	allDone.Add(3)

	go doCount(
		&allDone,
		[]chan counters.Payload{c12, c13},
		[]chan counters.Payload{c21, c31},
	)
	go doCount(
		&allDone,
		[]chan counters.Payload{c21, c23},
		[]chan counters.Payload{c12, c32},
	)
	go doCount(
		&allDone,
		[]chan counters.Payload{c31, c32},
		[]chan counters.Payload{c13, c23},
	)

	allDone.Wait()
}

func doCount(allDone *sync.WaitGroup, txs, rxs []chan counters.Payload) {
	defer allDone.Done()
	id := uuid.Must(uuid.NewV4())
	counter := counters.NewGCounter(id, nil)
	counter.Inc()
	counter.Inc()
	counter.Inc()

	payload := counter.Serialize()
	for _, tx := range txs {
		tx <- payload
	}

	for _, rx := range rxs {
		received := <-rx
		counter.Merge(received)
	}

	fmt.Printf("Hello from %s: %d\n", id, counter.Val())
}
