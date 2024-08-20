package main

import "fmt"

// Реализовать паттерн «адаптер» на любом примере.

type Dog struct{}

func (dog *Dog) WoofWoof() {
	fmt.Println("Gaf Gaf")
}

type Cat struct{}

func (cat *Cat) MeowMeow(isCall bool) {
	if isCall {
		fmt.Println("Mjau Mjau")
	}
}

type Wife struct {
}

func (wife *Wife) Reaction() {
	fmt.Println("WIFE")
}

type AnimalAdapter interface {
	Reaction()
}
type DogAdapter struct {
	*Dog
}

func (adapter *DogAdapter) Reaction() {
	adapter.WoofWoof()
}
func NewDogAdapter(dog *Dog) AnimalAdapter {
	return &DogAdapter{dog}
}

type CatAdapter struct {
	*Cat
}

func (adapter *CatAdapter) Reaction() {
	adapter.MeowMeow(true)
}
func NewCatAdapter(cat *Cat) AnimalAdapter {
	return &CatAdapter{cat}
}

func main() {

	fmt.Println("\nВы останавливаетесь перед дверью и вставляете в ухо адаптер с двумя чипами\n")
	myFamily := [3]AnimalAdapter{NewDogAdapter(&Dog{}), NewCatAdapter(&Cat{}), &Wife{}}
	//
	fmt.Println("Открываете дверь и заходите домой\n")
	for _, member := range myFamily {
		member.Reaction()

	}

}
