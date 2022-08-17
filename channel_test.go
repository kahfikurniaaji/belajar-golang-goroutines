package belajar_golang_goroutines

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	var channel chan string = make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Kahfi Kurnia Aji"
		fmt.Println("Selesai Mengirim Data ke Channel")
	}()

	var data string = <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Kahfi Kurnia Aji"
}

func TestChannelAsParameter(t *testing.T) {
	var channel chan string = make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	var data string = <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Kahfi Kurnia Aji"
}

func OnlyOut(channel <-chan string) {
	var data string = <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	var channel chan string = make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	var channel chan string = make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "Kahfi"
		channel <- "Kurnia"
		channel <- "Aji"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Selesai")
}

func TestRangeChannel(t *testing.T) {
	var channel chan string = make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Menerima data", data)
	}

	fmt.Println("Selesai")
}

func TestSelectChannel(t *testing.T) {
	var channel1 chan string = make(chan string)
	var channel2 chan string = make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari Channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari Channel 2", data)
			counter++
		}
		if counter == 2 {
			break
		}
	}
}

func TestDefaultSelectChannel(t *testing.T) {
	var channel1 chan string = make(chan string)
	var channel2 chan string = make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari Channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari Channel 2", data)
			counter++
		default:
			fmt.Println("Menunggu Data")
		}
		if counter == 2 {
			break
		}
	}
}
