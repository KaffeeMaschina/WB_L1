package main

import "fmt"

//Разработать конвейер чисел.
//Даны два канала: в первый пишутся числа (x) из массива,
//во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.

func main() {
	nums := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// First variant
	ch := prod(nums[:]...)
	out := cons(ch)
	for i := 0; i <= len(nums)-1; i++ {
		fmt.Println(<-out)
	}
	// Second variant
	for n := range cons(prod(nums[:]...)) {
		fmt.Println(n)
	}
}

// Write to first channal x from an array
func prod(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// Read from first channel and write to second channel x*2
func cons(in <-chan int) <-chan int {

	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * 2
		}
		close(out)
	}()
	return out
}
