package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// Реализовать постоянную запись данных в канал (главный поток).
//Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
//Необходима возможность выбора количества воркеров при старте.
//Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.

const ProgramStoppedByUser = 100

func work(ctx context.Context, id int, unbufferedChannel <-chan int64) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d is closed \n", id)
			return
		case message, ok := <-unbufferedChannel:
			if ok {
				fmt.Println("id", id, "message", message)
			}
		}
	}

}

func main() {
	// Выбор количества воркеров при старте.
	var numWorker int
	fmt.Scan(&numWorker)

	var wg sync.WaitGroup

	unbufferedChannel := make(chan int64)
	defer close(unbufferedChannel)

	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT)

	// Осуществляется постоянная запись случайных чисел в канал.
	go func() {
		for {
			x := rand.Int63n(100)
			unbufferedChannel <- x

		}
	}()
	// Воркеры читают произвольные данные из канала и выводят в stdout.
	for i := 0; i <= numWorker; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			work(ctx, i, unbufferedChannel)
		}(i)
	}
	wg.Wait()
	for {
		select {
		// Главная программа завершится: закроет канал и вернет код ответа.
		case <-ctx.Done():
			close(unbufferedChannel)
			log.Println("Stopped by user")
			os.Exit(ProgramStoppedByUser)

		}
	}
}
