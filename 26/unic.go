package main

import (
	"fmt"
	"strings"
)

// Разработать программу, которая проверяет, что все символы в строке уникальные
// (true — если уникальные, false etc).
// Функция проверки должна быть регистронезависимой.

// abcd — true
// abCdefAaf — false
// aabcd — false

// Берём значение, если оно не последнее, то сравниваем его со следующим.
func main() {
	s1 := "abCD"
	s2 := "abaAB"
	fmt.Println(s1, ifUnic(s1))
	fmt.Println(s2, ifUnic(s2))
}
func ifUnic(str string) bool {
	s := strings.ToLower(str)
	m := make(map[rune]int)
	var r bool = true

	for _, ch := range s {
		m[ch] += 1
		if m[ch] > 1 {
			r = false
		}
	}
	return r
}
