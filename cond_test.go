package golang_goroutine

/*
	Cond adalah adalah implementasi locking berbasis kondisi.
	Cond membutuhkan Locker (bisa menggunakan Mutex atau RWMutex) untuk implementasi locking nya, namun berbeda dengan Locker biasanya, di Cond terdapat function Wait() untuk menunggu apakah perlu menunggu atau tidak
	Function Signal() bisa digunakan untuk memberi tahu sebuah goroutine agar tidak perlu menunggu lagi, sedangkan function Broadcast() digunakan untuk memberi tahu semua goroutine agar tidak perlu menunggu lagi
	Untuk membuat Cond, kita bisa menggunakan function sync.NewCond(Locker)
*/

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var cond = sync.NewCond(&sync.Mutex{})
var group = sync.WaitGroup{}

func WaitCondition(value int) {
	defer group.Done()

	cond.L.Lock()
	cond.Wait()
	fmt.Println("Done", value)
	cond.L.Unlock()
}

func TestCond(t *testing.T) {
	for i := 0; i < 10; i++ {
		group.Add(1)
		go WaitCondition(i)
	}

	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		time.Sleep(1 * time.Second)
	// 		cond.Signal() // secara 1 1
	// 	}
	// }()

	go func() {
		time.Sleep(1 * time.Second)
		cond.Broadcast() // semuanya
	}()

	group.Wait()
	fmt.Println("SELESAI")
}
