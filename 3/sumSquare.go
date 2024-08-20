package main

import (
	"fmt"
	"time"
)

//Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.

func main() {

	go goSumSquare(2, 4, 6, 8, 10)
	time.Sleep(time.Second)
}
func goSumSquare(nums ...int) {
	total := 0
	for i := 0; i < len(nums); i++ {
		nums[i] *= nums[i]
	}
	fmt.Println(nums)
	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}
	fmt.Println(total)
}
