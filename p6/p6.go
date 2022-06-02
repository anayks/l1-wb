package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

const KILLS_TYPES = 3

/*
	На самом деле - тип завершения работы всего 1 - ожидание чтения из канала,
	Просто разными способами - через время, контекстом или напрямую обменом данными.

	Можно ещё сделать костыль через запись в канал, но это бессмысленно
*/

func main() {
	var wg sync.WaitGroup
	wg.Add(KILLS_TYPES)

	ctx := context.Background()
	ctx, cancelFunc := context.WithCancel(ctx)
	go GoroutineWithContext(&wg, ctx)
	cancelFunc()

	ch := make(chan struct{})
	go GoroutineWithChan(&wg, ch)
	ch <- struct{}{}

	go GoroutineWithTimer(&wg)
	wg.Wait()
}

func GoroutineWithContext(wg *sync.WaitGroup, ctx context.Context) { // ожидаем, пока контекст будет отменен, и когда это происходит - завершается работа
	defer wg.Done()
	<-ctx.Done()
	fmt.Printf("\nРабота горутины с контекстом закончена!")
}

func GoroutineWithChan(wg *sync.WaitGroup, ch chan struct{}) { // ожидаем запись в канал. Если что-то было написано, то завершается работа
	defer wg.Done()
	<-ch
	fmt.Printf("\nРаботы горутины с каналом закончена!")
}

func GoroutineWithTimer(wg *sync.WaitGroup) { // ждем время и по истечении завершается работа
	defer wg.Done()
	<-time.After(time.Second * 1)
	fmt.Printf("\nПрошла 1 секунда, горутина с таймером выключилась!")
}
