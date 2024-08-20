package main

import (
	"fmt"
)

//Разработать программу, которая переворачивает подаваемую на ход строку
//(например: «главрыба — абырвалг»). Символы могут быть unicode.

func main() {

	var newString string = "главрыба"

	ns := inverter([]rune(newString))
	fmt.Println(string(ns))
}

func inverter(s []rune) []rune {

	var new []rune

	for i := len(s) - 1; i >= 0; i-- {
		a := s[i]
		new = append(new, a)

	}
	return new
}
