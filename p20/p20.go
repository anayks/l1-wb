package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("%v", reverseWordsInString("привет пёс рыба последний"))
}

func reverseWordsInString(str string) string {
	result := strings.Split(str, " ")                      // разделяем слова (строки, между которыми пробел) и получаем массив строк
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 { // так же проходимся по массиву строку и меняем их местами
		result[i], result[j] = result[j], result[i]
	}
	return strings.Join(result, " ") // возвращаем результат в виде строки, которая получается проставлением пробелов между словами из массива
}
