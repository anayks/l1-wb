package main

import (
	"fmt"
	"math"
)

var data map[int][]float64

func main() {
	data = map[int][]float64{}
	groupArray([]float64{1, 2, -10, -15, -43.4, -40, -41.4, 3.10, 10.4, 20.4})

	fmt.Printf("data: %v", data)
}

func groupArray(arr []float64) {
	for _, v := range arr {
		var result float64
		if v >= 0 {
			result = math.Floor(v/10) * 10 // частное от 10 умножаем на 10, чтобы получить десятки без остатка
		} else { // в примере указано, что -27 относится к -20, поэтому делаем такой костыль, чтобы при отрицательном числе относилось к большему значению
			result = math.Ceil(v/10) * 10
		}
		_, ok := data[int(result)]
		if !ok { // если значения нет, создаем его и вставляем туда значение
			data[int(result)] = make([]float64, 0)
			data[int(result)] = append(data[int(result)], v)
			continue
		}
		data[int(result)] = append(data[int(result)], v) // если значение есть, вставляем в него значение
	}
}
