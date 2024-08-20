package main

import (
	"fmt"
	"math/big"
)

// Разработать программу, которая перемножает, делит, складывает,
// вычитает две числовых переменных a,b, значение которых > 2^20.

func main() {
	withBigNumbers()
	withInt64()
}

// Если подразумевалось использовать очень большие числа
func withBigNumbers() {
	a := big.NewInt(13980374560941)
	b := big.NewInt(12123452345324)

	add := new(big.Int)
	fmt.Printf("%v + %v = %v\n", a, b, add.Add(a, b))

	sub := new(big.Int)
	fmt.Printf("%v - %v = %v\n", a, b, sub.Sub(a, b))

	mul := new(big.Int)
	fmt.Printf("%v * %v = %v\n", a, b, mul.Mul(a, b))

	div := new(big.Int)
	fmt.Printf("%v / %v = %v\n", a, b, div.Div(a, b))
}

func withInt64() {
	// 64 бита вмещают результаты делений, перемножений и т.д. чисел > 2^20
	// 2^20*2^20 имеет 13 знаков, а int64 - 18 знаков
	var a, b int64 = 123564574457, 43563234546

	fmt.Printf("a*b = %v\n", a*b)
	fmt.Printf("a/b = %v\n", a/b)
	fmt.Printf("a-b = %v\n", a-b)
	fmt.Printf("a+b = %v\n", a+b)
}
