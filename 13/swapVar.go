package main

import (
	"fmt"
)

// Поменять местами два числа без создания временной переменной.

func main() {
	var i, j = 1, 6
	fmt.Printf("i = %v, j = %v\n", i, j)
	i, j = j, i
	fmt.Printf("i = %v, j = %v\n", i, j)
}
