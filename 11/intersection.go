package main

import "fmt"

//Реализовать пересечение двух неупорядоченных множеств

func main() {

	slice := []int{116, 12, 12, 14}
	slice2 := []int{12, 12, 13, 12}

	fmt.Println(newSlice(slice, slice2))
}

func newSlice(s ...[]int) []int {
	ns := make([]int, 0)

	for _, sl := range s {
		for _, nv := range sl {
			if contains(ns, nv) == false {
				ns = append(ns, nv)
			}
		}
	}
	return ns
}
func contains(sl []int, s int) bool {
	for _, v := range sl {
		if v == s {
			return true
		}
	}
	return false

}
