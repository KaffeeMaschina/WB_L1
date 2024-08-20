package main

import (
	"fmt"
	"os"
	"time"
)

//Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать.
//По истечению N секунд программа должна завершаться.

func main() {
	unbufferedChannel := make(chan int)
	defer close(unbufferedChannel)

	timer2 := time.NewTimer(time.Second * 10)

	go writeToChan(unbufferedChannel)

	// First variant
	forSelect(*timer2, unbufferedChannel)

	//Ssecond variant
	/*go timer(*timer2)
	goroutineAndRange(unbufferedChannel)*/

}

func writeToChan(unbufferedChannel chan int) {
	for i := 0; ; i++ {
		unbufferedChannel <- i
		time.Sleep(time.Second)
	}
}

// First variant
func forSelect(timer2 time.Timer, unbufferedChannel chan int) {
	for {
		select {
		case <-timer2.C:
			fmt.Println("Timer 2 expired")
			os.Exit(1)
		default:
			v := <-unbufferedChannel
			fmt.Println(v)
		}

	}
}

// Second variant
func goroutineAndRange(unbufferedChannel chan int) {
	for v := range unbufferedChannel {
		fmt.Println(v)
	}
}
func timer(timer2 time.Timer) {
	<-timer2.C
	fmt.Println("Timer 2 expired")
	os.Exit(1)
}
