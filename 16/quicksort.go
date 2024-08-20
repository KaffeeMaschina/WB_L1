package main

import "fmt"

//Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

func main() {
	arr := []int{15, 33, 10, 12}
	quicksort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
func quicksort(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		quicksort(arr, low, p-1)
		quicksort(arr, p+1, high)
	}
	return arr
}

func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			fmt.Println(i, j)
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}
