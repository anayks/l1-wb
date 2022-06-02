package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	inc1 := NewIncrementator1()
	inc2 := NewIncrementator2()

	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			inc1.Add(1)
		}()
		inc1.Add(1)
	}

	wg.Wait()

	fmt.Printf("\nvalue of incrementator1: %v", inc1.val) // 6
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			inc2.Add(3) // 3 * 3 = 9
		}()
		inc2.Add(2) // 2 * 3 = 6
	}

	wg.Wait()

	fmt.Printf("\nvalue of incrementator2: %v", inc2.val) // 15
}

type Incrementator1 struct { // через мьютекс
	sync.Mutex
	val int
}

func NewIncrementator1() *Incrementator1 {
	return &Incrementator1{}
}

func (i *Incrementator1) Add(count int) {
	i.Lock()
	defer i.Unlock()
	i.val += count
}

type Incrementator2 struct {
	val int64
}

func NewIncrementator2() *Incrementator2 {
	return &Incrementator2{}
}

func (i *Incrementator2) Add(count int64) {
	atomic.AddInt64(&i.val, count)
}
