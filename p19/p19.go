package main

import "fmt"

func main() {
	fmt.Printf("Result: %v", reverseString("привет, как дела?12ac1"))
}

func reverseString(val string) (result string) {
	runes := []rune(val)                                  // создаем слайс рун
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 { // в цикле проходимся слева и справа одновременно с двумя переменными. Когда меньшая станет больше, чем большая
		runes[i], runes[j] = runes[j], runes[i] // меняем значения местами
	}
	result = string(runes) // слайс рун переводим в строку
	return result          // возвращаем строку
}
