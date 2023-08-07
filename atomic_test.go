package golang_goroutine

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

/*
	Atomic merupakan package yang digunakan untuk menggunakan data primitive secara aman pada proses concurrent
	Contohnya sebelumnya kita telah menggunakan Mutex untuk melakukan locking ketika ingin menaikkan angka di counter. Hal ini sebenarnya bisa digunakan menggunakan Atomic package
*/

func TestAtomic(t *testing.T) {
	var x int64 = 0
	group := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		go func() {
			defer group.Done()
			group.Add(1)
			for j := 0; j < 100; j++ {
				atomic.AddInt64(&x, 1)
			}
		}()
	}

	group.Wait()
	fmt.Println("Result:", x)
}
