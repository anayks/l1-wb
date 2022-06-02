package main

import (
	// "fmt"
	"math"
)

type Point struct {
	x, y float64 // значение инкапсулированы в пакете, так как называются с маленькой буквы
}

func NewPoint(x, y float64) Point { // Фабрика точек
	return Point{
		x: x,
		y: y,
	}
}

func (p Point) Dist(t Point) float64 { // Функция не инкапсулирована, но может работать с данными, обозначенными в этом пакете
	return math.Sqrt(math.Pow(p.x-t.x, 2) + math.Pow(p.y-t.y, 2))
}

// func main() {
// 	p1 := NewPoint(3, 4)
// 	p2 := NewPoint(0, 0)
// 	fmt.Printf("Distance: %v", p1.Dist(p2))
// }
