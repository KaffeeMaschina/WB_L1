package main

import "fmt"

//Дана структура Human (с произвольным набором полей и методов).
//Реализовать встраивание методов в структуре Action
//от родительской структуры Human (аналог наследования).

type Human struct {
	Name string
	Age  int
	a    Action
}
type Action struct {
}

func main() {
	h := Human{
		Name: "Bob",
		Age:  24,
	}
	h.a.Move(h)
	h.a.Run()
}
func (a Action) Move(h Human) {
	fmt.Printf("Moving human, %v, %v\n", h.Name, h.Age)
}
func (a Action) Run() {
	fmt.Println("Running human")
}
