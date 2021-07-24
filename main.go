package main

import (
	"fmt"
	"sync"
)

var m = new(sync.Mutex)
var m2 = new(sync.Mutex)

func main() {
	var table = make(map[int]struct{})
	var table2 = make(map[int]struct{})
	go func() {
		for {
			m.Lock()
			table[0] = struct{}{}
			table2[0] = struct{}{}
			m.Unlock()
		}
	}()
	go func() {
		for {
			m.Lock()
			fmt.Println("1", table[0])
			fmt.Println("2", table2[0])
			m.Unlock()
		}
	}()
	done := make(chan struct{})
	defer close(done)
	select {
	case _ = <-done:
	}
}
