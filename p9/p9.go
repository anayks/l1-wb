package main

import "fmt"

func main() {
	ch := createConveyor([]int{1, 3, 7, 20}) // должно вернуть канал с квадратами чисел: 1, 9, 49, 400

	for v := range ch { // в цикле читаем данные из канала, в котором выгружены квадраты чисел
		fmt.Printf("Data: %v\n", v)
	}
}

func createConveyor(arr []int) chan int {
	ch := genChan(arr)
	return updateChan(ch)
}

func genChan(arr []int) chan int { // функция создает канал для обработки чисел и загружает в него числа из массива
	ch := make(chan int)
	go func() {
		for _, v := range arr {
			ch <- v
		}
		close(ch)
	}()
	return ch
}

func updateChan(gen chan int) chan int { // функция читает данные из предыдущего канала и пишет в следующий значения в квадрате
	ch := make(chan int)
	go func() {
		for v := range gen {
			ch <- v * v
		}
		close(ch)
	}()
	return ch
}
