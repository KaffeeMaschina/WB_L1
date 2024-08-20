package main

import (
	"fmt"
	"time"
)

// Реализовать собственную функцию sleep.

func main() {
	fmt.Println("before sleep")
	sleep(5 * time.Second)
	fmt.Println("after sleep")
}
func sleep(d time.Duration) {
	<-time.After(d)
}
