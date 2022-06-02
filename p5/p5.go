package main

import (
	"fmt"
	"math/rand"
	"time"
)

const TIME_IN_SECONDS = 5

func main() {
	ch := make(chan int, 1) // то что здесь будет буфер в 1 элемент, гарантирует то, что данные буквально из for ниже будут отправлены по очереди

	go func() {
		for {
			fmt.Printf("\nЯ последовательно читаю: %v", <-ch)
		}
	}()

	for {
		select {
		case <-time.After(time.Second * TIME_IN_SECONDS):
			{
				fmt.Printf("\nВремя истекло.")
			}
		default:
			{
				time.Sleep(time.Millisecond * 300)
				ch <- rand.Int() // отправка в буферизированный канал. Если буфер переполнен - горутина ждет и отправляет, когда канал освободится.
			}
		}
	}
}
