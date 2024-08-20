package main

import (
	"fmt"
	"math"
)

//Разработать программу нахождения расстояния между двумя точками,
//которые представлены в виде структуры Point
//с инкапсулированными параметрами x,y и конструктором.

type point struct {
	x, y int
}

func main() {

	a := newPoint(1, 12)
	b := newPoint(1, 13)

	fmt.Println(distance(a, b))

}

func newPoint(x, y int) point {
	return point{
		x: x,
		y: y,
	}
}

func distance(a, b point) float64 {
	return math.Sqrt(math.Pow(float64(a.x-b.x), 2) + math.Pow(float64(a.y-b.y), 2))
}
