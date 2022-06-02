package main

import "sync"

// go run -race p7/p7.go

var dataMapV1 map[string]int
var mx sync.Mutex

func main() {
	dataMapV1 = map[string]int{}
	go func() {
		WriteInMap("test1", 35)
	}()
	WriteInMap("test1", 3)

	m2 := NewMapV2()
	go func() {
		m2.WriteInMap("test2", 35)
	}()
	m2.WriteInMap("test2", 34)
}

func WriteInMap(key string, data int) { // просто используем мьютекс и обращаемся к данным безопасно
	mx.Lock()
	defer mx.Unlock()

	dataMapV1[key] = data
}

type MapStructV2 struct {
	data map[string]int
	sync.Mutex
}

func NewMapV2() MapStructV2 {
	return MapStructV2{
		data: map[string]int{},
	}
}

func (m *MapStructV2) WriteInMap(key string, data int) { // Можем использовать для этого кастомную структуру и мьютекс хранить связно с мапой
	m.Lock()
	defer m.Unlock()

	m.data[key] = data
}
