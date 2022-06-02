package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Printf("Data: %v %v %v", isSymbolsUnique("test"), isSymbolsUnique("tes"), isSymbolsUnique("1Afga")) // false, true, false
}

func isSymbolsUnique(text string) bool {
	runes := []rune(text)             // создаем слайс рун из строк
	symbols := make(map[rune]bool, 0) // создаем мапу из уникальных символов

	for _, v := range runes { // в цикле проходимся по слайсу рун
		v = unicode.ToLower(v) // делаем руну маленькой
		_, ok := symbols[v]
		if !ok { // если руна не была ни разу записана в мапу
			symbols[v] = true // записываем значение true
			continue
		}
		return false // если руна хоть раз была записана в мапу, то строка не уникальна, возвращаем false
	}
	return true // прошлись по всем символам, совпадений нет, символы в строке уникальны
}
