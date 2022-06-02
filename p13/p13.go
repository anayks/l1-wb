package main

import "fmt"

func main() {
	var a, b int = 1, 3
	SwapElements1(&a, &b)
	fmt.Printf("a: %v, b: %v\n", a, b) // 3, 1
	SwapElements2(&a, &b)
	fmt.Printf("a: %v, b: %v\n", a, b) // 1, 3
	SwapElements3(&a, &b)
	fmt.Printf("a: %v, b: %v\n", a, b) // 3, 1
}

func SwapElements1(a *int, b *int) { // свойствами языка
	*a, *b = *b, *a
}

func SwapElements2(a *int, b *int) { // вычитанием
	*a = *a + *b
	*b = *a - *b
	*a = *a - *b
}

func SwapElements3(a *int, b *int) { // XOR
	*a = *a ^ *b
	*b = *a ^ *b
	*a = *a ^ *b
}
