package main

import "fmt"

func main() {
	fmt.Printf("Result: %v", GetIntersection([]int{1, 1, 3, 25, 64, 535, 222}, []int{1, 1, 64, 33333, 3535, 535, 22222, 222}))
}

func GetIntersection(arr1, arr2 []int) (result []int) {
	var uniqValues map[int]bool // создаем мапу уникальных значений
	uniqValues = map[int]bool{}

	for _, v := range arr1 { // если сталкиваемся с каким-то значением, записываем его в мапу, что столкнулись с ним
		uniqValues[v] = true
	}

	for _, v := range arr2 {
		_, ok := uniqValues[v] // если значение не было записано - пропускаем
		if !ok {
			continue
		}
		result = append(result, v) // записываем значение в результат
	}
	return result // возвращаем результат
}
