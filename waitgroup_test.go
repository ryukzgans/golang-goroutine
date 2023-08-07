package golang_goroutine

// WaitGroup adalah fitur yang bisa digunakan untuk menunggu sebuah proses selesai dilakukan

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsynchronous(group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)

	fmt.Println("Test")
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	counter := 0
	for i := 0; i < 100; i++ {
		counter++
		go RunAsynchronous(group)
	}

	group.Wait()
	fmt.Println("Selesai")
	fmt.Println(counter)
}
