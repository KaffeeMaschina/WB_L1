package main

import (
	"errors"
	"fmt"
)

//Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

type set struct {
	elements map[string]struct{}
}

func main() {
	set := createSet()

	set.add("cat")
	set.add("cat")
	set.add("dog")
	set.add("cat")
	set.add("tree")

	set.list()
}
func createSet() *set {
	var set set
	set.elements = make(map[string]struct{})
	return &set
}
func (set *set) add(s string) {
	set.elements[s] = struct{}{}
}
func (set *set) delete(s string) error {
	if _, exists := set.elements[s]; !exists {
		return errors.New("string doesn't exist")
	}
	delete(set.elements, s)
	return nil
}
func (set *set) contains(s string) bool {
	_, exists := set.elements[s]
	return exists
}
func (set *set) list() {
	for key := range set.elements {
		fmt.Println(key)
	}
}
