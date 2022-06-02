package main

import (
	"fmt"
	"reflect"
)

type DataInterface interface{}

func main() {
	var d interface{} = make(chan int)
	switch d.(type) {
	case int:
		{
			fmt.Println("int")
		}
	case bool:
		{
			fmt.Println("bool")
		}
	case string:
		{
			fmt.Println("string")
		}
	case chan interface{}: // не пройдет, потому что тип канала не interface{}
		{
			fmt.Println("channel")
		}
	case chan int: // нужно явно указывать тип канал
		{
			fmt.Println("channel")
		}
	default:
		{
			fmt.Println("unknown")
		}
	}
	main1()
}

func main1() { // вариант два, можно явно вернуть тип элемента
	var d interface{} = make(chan int)
	xType := reflect.TypeOf(d)
	fmt.Printf("Type of second d: %v", xType)
}

// Можно так же распарсить строку, которая возвращается рефлектом, и, если там есть chan, то возвращать channel соответственно
