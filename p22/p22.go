package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	firstValue := math.Pow(2, 20)
	secondValue := math.Pow(2, 25)
	result := multiplyBigInt(int64(firstValue), int64(secondValue))
	fmt.Printf("Result of multiply 2^20 * 2^25: %v", result.String()) // 35 184 372 088 832 == 2^45
	thirdValue := math.Pow(2, 35)
	result = divBigInt(int64(result.Int64()), int64(thirdValue))
	fmt.Printf("\nResult of div 2^45 on 2^35: %v", result.String()) // 1 024 == 2^10
	fourthValue := math.Pow(2, 35)
	result = differenceBigInt(int64(result.Int64()), int64(fourthValue))
	fmt.Printf("\nResult of div 2^45 on 2^35: %v", result.String()) // -34 359 737 344
	fifthValue := math.Pow(2, 35)
	result = addBigInt(int64(result.Int64()), int64(fifthValue))
	fmt.Printf("\nResult of div 2^45 on 2^35: %v", result.String()) // 1 024
}

func multiplyBigInt(a, b int64) (c big.Int) { // умножение больших чисел
	firstValue := big.NewInt(a)
	secondValue := big.NewInt(b)
	return *c.Mul(firstValue, secondValue)
}

func differenceBigInt(a, b int64) (c big.Int) { // вычитание больших чисел
	firstValue := big.NewInt(a)
	secondValue := big.NewInt(b)
	return *c.Sub(firstValue, secondValue)
}

func addBigInt(a, b int64) (c big.Int) { // сложение больших чисел
	firstValue := big.NewInt(a)
	secondValue := big.NewInt(b)
	return *c.Add(firstValue, secondValue)
}

func divBigInt(a, b int64) (c big.Int) { // деление больших чисел нацело
	firstValue := big.NewInt(a)
	secondValue := big.NewInt(b)
	return *c.Div(firstValue, secondValue)
}
