package main

import (
	"fmt"
	"log"
	"strconv"
)

//Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

func changeOneBit(num, position, newBit int) int64 {
	n := strconv.FormatInt(int64(num), 2)
	fmt.Println(n)
	bits := []byte(n)

	nb := strconv.FormatInt(int64(newBit), 2)
	newB := []byte(nb)
	fmt.Println(nb)

	bits[position] = newB[0]
	newVal, err := strconv.ParseInt(string(bits), 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	return newVal
}
func main() {
	fmt.Println(changeOneBit(5, 1, 1)) //101 = 5 -> 111 = 7
}
