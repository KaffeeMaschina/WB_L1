package main

import "fmt"

// Удалить i-ый элемент из слайса
func main() {
	slice := []int{-5, -4, 0, -11, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(delete(9, slice))
}

func delete(i int, sl []int) []int {
	var newSlice []int = sl[0:i]
	var newSLice2 []int = sl[i+1:]
	return append(newSlice, newSLice2...)
}
