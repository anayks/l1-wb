package main

type Human struct {
}

func (h Human) Speak() {}

func (h Human) Eat() {}

type Action struct {
	Human
}

func main() {
	act := Action{
		Human{},
	}
	act.Eat()
}

// Готово? Базовый функционал языка Go по внедрению методов одной структуры из другой
