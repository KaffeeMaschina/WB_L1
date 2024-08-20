package main

import (
	"fmt"
	"strings"
)

//Разработать программу, которая переворачивает слова в строке.
//Пример: «snow dog sun — sun dog snow».

func main() {
	var s string

	s = "snow dog sun"
	wordsInverterOnlyEnglish(s)
	wordsInverter(s)
}
func wordsInverterOnlyEnglish(s string) {
	var sl, newSl []string
	var ns string
	var p int

	for i, r := range s {
		lit := string(r)

		if lit == " " {

			if p == 0 {
				sl = append(sl, s[:i+1])
			} else {
				sl = append(sl, s[p:i+1])
			}
			p = i + 1
		}
		if i == len(s)-1 {
			lastLit := s[p:i+1] + " "
			sl = append(sl, lastLit)
		}
	}
	for i := len(sl) - 1; i >= 0; i-- {
		newSl = append(newSl, sl[i])
	}
	for i := range newSl {
		ns += newSl[i]
	}
	fmt.Println(ns)
}
func wordsInverter(s string) {
	var ns []string

	newSl := strings.Fields(s)

	for i := len(newSl) - 1; i >= 0; i-- {

		ns = append(ns, newSl[i])
		s = strings.Join(ns, " ")

	}
	fmt.Println(s)
}
