package golang_goroutine

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

/*
	Membuat channel
	channel := make(chan typedata) >> make(chan string)
*/

/*
	Mengirim dan Menerima Data dari Channel
	channel <- "Ilham Kurniawan"  // ini contoh mengirim data ke channel
	data := <- channel 			 // ini contoh menerima data dari channel
*/

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel) // disarankan untuk selalu mengclose sebuah channel, dgn menggunakan defer baik error maupun tidak akan selalu dijalankan

	// anonymous function
	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Ilham Kurniawan Tarmiji" // mengirim data string ke channel
		fmt.Println("DATA BERHASIL DIKIRIM") // akan dijalankan setelah ada yang mengambil data dari channel, jika tidak akan terjadi block / deadlock
	}() // lgsung di jalankan ()

	data := <-channel // mengambil data dari channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

// # PARAMETER CHANNEL

// set parameter c sebagai channel bertipe data string
func GiveMeResponse(c chan string) {
	time.Sleep(2 * time.Second)
	c <- "ILHAM KURNIAWAN"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

// # Channel IN dan OUT

// hanya untuk mengirim data
func OnlyIn(c chan<- string) {
	time.Sleep(2 * time.Second)
	c <- "Ilham Kurniawan"
}

// hanya untuk menerima data
func OnlyOut(c <-chan string) {
	data := <-c
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

// #Buffered Channel (channel dgn x data)
// cap(channel) >> capacity (melihat panjang buffer)
// len(channel) >> length (melihat jumlah data di buffer)
func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3) // data buffer channel dgn maks 3 data
	defer close(channel)

	go func() {
		channel <- "Ilham"
		channel <- "Kurniawan"
		channel <- "Tarmiji"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("SELESAI")

}

// # Range Channel
func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke-" + strconv.Itoa(i)
		}
		close(channel) // dont forget to close the channel or it will be deadlock
	}()

	for data := range channel {
		fmt.Println(data)
	}

}

// # Select Channel
func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari Channel 1:", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari Channel 2:", data)
			counter++
		}
		if counter == 2 {
			break
		}

	}
}
