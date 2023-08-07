package golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

// menggunakan go, diawal pemanggilan function, tetapi dgn go kita tidak dpt menggunakan function return value
func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("END")

	time.Sleep(2 * time.Second)
}

func DisplayNumber(num int) {
	fmt.Println("Display", num)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(5 * time.Second)
}
