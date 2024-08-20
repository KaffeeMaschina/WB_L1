package main

import "fmt"

// Реализовать бинарный поиск встроенными методами языка.
func main() {
	arr := [100]int{}
	for i := 0; i <= len(arr)-1; i++ {
		arr[i] = i
	}
	fmt.Println(linearSearch(49, arr[:]))
	fmt.Println(binarySearch(50, arr[:]))
}
func linearSearch(n int, arr []int) (int, int) {
	var operations int
	for i, v := range arr {
		operations = i
		if n == v {
			return v, operations
		}
	}
	return -1, operations
}
func binarySearch(n int, arr []int) (int, int) {
	var l, r int = arr[0], len(arr) - 1
	var operations int

	for l <= r {
		operations++
		var m int = (l + r) / 2
		mv := arr[m]
		if n == m {
			return mv, operations
		} else if m < n {
			l = m + 1
		} else if m > n {
			r = m - 1
		}

	}
	return -1, operations
}
