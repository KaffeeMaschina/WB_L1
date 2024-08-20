package main

import "fmt"

//Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
//Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.
//Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.

func main() {
	arr := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	fmt.Println(repeatSelect(arr))

}
func repeatSelect(arr []float32) map[int][]float32 {
	m := make(map[int][]float32)

	for i := -10; i >= -40; i = i - 10 {
		m[i] = selec(float32(i), arr)
	}

	for i := 0; i <= 40; i = i + 10 {
		m[i] = selec(float32(i), arr)
	}
	return m
}
func selec(min float32, arr []float32) (arrb []float32) {
	var max float32
	if min < 0 {
		max = min - 10.0
		for _, v := range arr {
			if min >= v && v >= max {
				arrb = append(arrb, v)
			}
		}
	} else {
		max = min + 10.0
		for _, v := range arr {
			if min <= v && v <= max {
				arrb = append(arrb, v)
			}
		}
	}
	return
}
