package main

import "fmt"

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Printf("Delete element: %v", deleteElement(3, arr))
}

func deleteElement(n int, arr []int) []int {
	if n < 0 { // если нужно удалить номер меньше 0, возвращаем тот же слайс
		return arr
	}

	if n > len(arr)-1 { // если номер больше номера последнего элемента, возвращаем тот же массив
		return arr
	}

	if n == 0 { // если номер равен 0, возвращаем все от нуля
		return arr[1:]
	}

	if n == len(arr)-1 { // если номер - это последний номер элемента в слайсе, возвращаем слайс кроме него
		return arr[:len(arr)-1]
	}

	newArr := make([]int, len(arr)-1) // создаем слайс, который меньше размером на 1 элемент
	copy(newArr, arr[:n])             // копируем в этот слайс всё до этого номера
	copy(newArr[n:], arr[n+1:])       // копируем в слайс после этого номера всё то, что после этого номера до конца
	return newArr                     // возвращаем новое значение
}
