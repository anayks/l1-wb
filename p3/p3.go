package main

import (
	"fmt"
	"math"
)

func main() {
	GetAllDoublePow([]int{2, 4, 6, 8, 10})
}

func GetResultFromChan(ch chan int, quitChan chan struct{}) int {
	var summ int

	for {
		select {
		case res := <-ch:
			{
				summ += res
			}
		case <-quitChan:
			{
				return summ
			}
		}
	}

	// здесь используем бесконечный цикл for, чтобы в select считывать данные с двух каналов
	// оно будет каждый раз срабатывать на первом канале, пока там не закончатся данные
	// и когда данные закончатся, оно прочитает данные с канала выхода и выйдет из функции с результатом решения
}

func GetAllDoublePow(arr []int) {
	ch := make(chan int)
	quitChan := make(chan struct{}, 1)

	go func(arrData []int, quitChan chan struct{}) {
		for _, v := range arrData {
			res := math.Pow(float64(v), 2)
			ch <- int(res)
		}
		quitChan <- struct{}{}
	}(arr, quitChan)

	// конкурентно запускаем функцию, которая в цикле считает квадраты и передает их в канал
	// по завершении цикла записываем в канал выхода, что запись завершено

	fmt.Printf("summ: %v", GetResultFromChan(ch, quitChan))
}
