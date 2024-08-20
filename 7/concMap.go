package main

import (
	"fmt"
	"sync"
)

// Реализовать конкурентную запись данных в map.
func main() {
	m := make(map[int]int)
	var mu sync.RWMutex

	for i := 0; i <= 100; i++ {
		go func(i int) {
			mu.Lock()
			defer mu.Unlock()
			m[i+2] = i
		}(i)
	}

	fmt.Println(m)
}
