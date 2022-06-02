package main

import "fmt"

type AliveEntity interface { // интерфейс живой сущности, которая должна уметь разговаривать
	Speak(text string)
}

type Alives []AliveEntity // Слайс живых сущностей

type Human struct{} // структура человека, он может только шептать, не говорить

func main() {
	aliveEntities := make(Alives, 0)                       // инициализируем слайс живых сущностей
	aliveEntities = append(aliveEntities, HumanAdapter1{}) // вставляем туда адаптер человека, который может только шептать, но шепот будет считаться разговором
}

func (h *Human) Whisper(text string) { // функция шепота человека
	fmt.Printf("\nHuman whisping: %v", text)
}

type HumanAdapter1 struct { // адаптер через встраивание методов
	Human
}

func (a HumanAdapter1) Speak(text string) { // функция, реализующая разговор адаптера через шепот человека
	a.Whisper(text)
}
