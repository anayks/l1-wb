package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const WORKERS_COUNT = 25

func main() {
	ch := make(chan interface{}) // создаем канал с произвольными данными (пустой интерфейс)

	var wg sync.WaitGroup

	ctx := context.Background()
	ctx, cancelFunc := context.WithCancel(ctx) // создаем контекст с отменой, чтобы потом ловить эту отмену в других горутинах

	StartWorkers(WORKERS_COUNT, ch, ctx, &wg) // Запускаем воркеры

	osChan := make(chan os.Signal)
	signal.Notify(osChan, syscall.SIGINT, syscall.SIGTERM)

	WriteRandomData(&wg, osChan, ch) // основная горутина горутина с записью

	cancelFunc() // дождались сигнал, закрываем контекст
	wg.Wait()    // ждем завершения работы всех горутин

	fmt.Printf("\nAll workers ended work")
}

func WriteRandomData(wg *sync.WaitGroup, osChan chan os.Signal, ch chan interface{}) {
	wg.Add(1)
	defer wg.Done() // завершаем работу в ожидании одной работы
	for {
		select {
		case <-osChan: // ловим сигнал системы и завершаем работу функции
			{
				fmt.Printf("\nWorker writer ended work...")
				return
			}
		case <-time.After(time.Second): // раз в секунду пишем произвольные данные
			{
				rand.Seed(time.Now().Unix())
				typeRand := rand.Intn(5) // получаем случайное число от 0 до 5 и отправляем случайные данные в канал
				if typeRand == 0 {       // отправляем число
					rand := rand.Int()
					ch <- rand
				} else if typeRand == 1 { // отправляем структуру
					randNameNumber := rand.Int()
					randData := rand.Intn(10)
					type randomStruct struct {
						name string
						data int
					}

					data := randomStruct{}
					data.name = fmt.Sprintf("Игорь #%d", randNameNumber)
					data.data = randData
					ch <- data
				} else if typeRand == 2 { // отправляем массив
					randLength := 1 + rand.Intn(4)
					var data []int64
					for i := 0; i < randLength; i++ {
						randNumber := rand.Int()
						data = append(data, int64(randNumber))
					}
					ch <- data
				} else if typeRand == 3 { // отправляем строку
					data := fmt.Sprintf("Heyo, test string: %v", rand.Float64())
					ch <- data
				} else if typeRand == 4 { // отправляем float64
					data := rand.Float64()
					ch <- data
				}
			}
		}
	}
}

func StartWorkers(count int, ch chan interface{}, ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(WORKERS_COUNT) // добавляем, сколько нам нужно ждать горутин
	for i := 0; i < count; i++ {
		go func(data int) {
			defer wg.Done() // по завершении функции выполняем Done
			for {
				select {
				case <-ctx.Done(): // если контекст закрыт, выходим из функции
					{
						fmt.Printf("\nWorker listener #%v ended work", data)
						return
					}
				case result := <-ch: // читаем данные из канала ch
					{
						fmt.Printf("\ndata: %v, goroutine number: %v", result, data)
					}
				}
			}
		}(i)
	}
}
