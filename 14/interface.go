package main

import (
	"fmt"
	"reflect"
)

//Разработать программу, которая в рантайме способна определить
//тип переменной: int, string, bool, channel из переменной типа interface{}.

func main() {
	//c := make(chan int)
	var v interface{} = "sdf"
	typeAssertion(v)
	withTypeSwitch(v)
	withReflectint(v)
}
func withReflectint(v interface{}) {
	vType := reflect.TypeOf(v)
	fmt.Println("With reflection: ", vType)
}
func withTypeSwitch(v interface{}) {
	fmt.Println("Type Switch")
	switch i := v.(type) {
	case int:
		fmt.Println("Int ", i)
	case string:
		fmt.Println("String ", i)
	case bool:
		fmt.Println("Bool: ", i)
	case chan int:
		fmt.Println("Chan int: ", i)
	}
}
func typeAssertion(v interface{}) {
	i, ok := v.(int)
	if ok {
		fmt.Println("Int ", i)
	}

	s, ok := v.(string)
	if ok {
		fmt.Println("String ", s)
	}

	b, ok := v.(bool)
	if ok {
		fmt.Println("Bool: ", b)
	}

	ch, ok := v.(chan int)
	if ok {
		fmt.Println("Chan int: ", ch)
	}
}
