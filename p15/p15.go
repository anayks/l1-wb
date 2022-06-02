package main

func main() {
	someFunc()
}

var justString string

func createHugeString(r int) string {
	return "smth213235q12ёйцауцпавр"
}

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100] // возможно проблема в обрезании по байтам, если там есть символы UTF-8 на русском языке, например, которые занимают 2 байта?
}

func someFuncFixed(limit int) string {
	v := createHugeString(1 << 10)
	runes := []rune(v)       // превращаем строку в массив рун (руны не зависят от байт)
	if len(runes) >= limit { // если количество рун больше, чем лимит
		return string(runes[:limit]) // то обрезаем массив рун по лимиту символов
	}
	return v // иначе возвращаем строку
}
