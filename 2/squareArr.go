package main

import (
	"fmt"
	"sync"
	"time"
)

//Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10)
//и выведет их квадраты в stdout.

// 1. 'time.Sleep(time.Second)' after 'go goSquareArr(a)'
// 2. 'runtime.GOMAXPROCS(1)' before 'go goSquareArr(a)' and 'runtime.Gosched()' after 'go goSquareArr(a)'
// 3. With sync.WaitGroup
func main() {
	a := [...]int{2, 4, 6, 8, 10}

	//runtime.GOMAXPROCS(1)

	go goSquareArr(a)
	time.Sleep(time.Second)

	//runtime.Gosched()

	//WaitGroupSquareArr(a)
}

func goSquareArr(a [5]int) {

	for i := 0; i < len(a); i++ {
		a[i] *= a[i]

	}

	fmt.Println(a)
}
func WaitGroupSquareArr(a [5]int) {
	var wg sync.WaitGroup
	for i := 0; i < len(a); i++ {
		wg.Add(1)
		go func(i int) {

			a[i] *= a[i]

			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(a)
}
