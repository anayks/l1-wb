package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	asyncDoublePow([]int{2, 4, 6, 8, 10}) // 4, 16, 36, 64, 100
}

func asyncDoublePow(arr []int) {
	length := len(arr)
	var wg sync.WaitGroup // создаем группу ожидания
	wg.Add(length)        // количество элементов, которые ждем - равно длине массива

	for _, v := range arr { // перебираем массив
		go func(val int) { // передаем значение явно аргументом функции
			defer wg.Done()                      // отложенное выполнение задачи в группе
			doubled := math.Pow(float64(val), 2) // возводим в 2 степень
			fmt.Printf("\n%v", doubled)          // выводим в консоль
		}(v) // передаем значение явно аргументом функции
	}
	// Выполнение, очевидно, будет в случайном порядке, так как функции выполняются в разных горутинах
	wg.Wait()
}
