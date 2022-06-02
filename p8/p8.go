package main

import (
	"fmt"
)

type bitint int64

func main() {
	b := bitint(0)
	b.setIntBit(1, 1)
	b.setIntBit(1, 2)
	b.setIntBit(0, 1)
	b.setIntBit(1, 3)
	b.setIntBit(0, 3)
	fmt.Printf("result: %v", b) // 4
}

func (b *bitint) setIntBit(data int, pos uint) {
	if data == 0 { // если хотим поставить 0, то вызываем функцию очистки
		b.clearBit(pos)
	} else if data == 1 { // если хотим поставить 1, то вызываем функцию установки бита
		b.setBit(pos)
	}
}

func (b *bitint) clearBit(pos uint) {
	testData := int(*b) // действие номер 3 работает только с int
	mask := ^(1 << pos)
	/* сдвигаем 1 на pos шагов влево и вызываем "XOR" - то есть, поразрядное исключающее или (крышечку).
	Проходимся по всем битам и если конкретно один бит равен 1, то оставляем, иначе 0   */
	testData &= mask      // поразрядное умножение. Если оба 1 - то 1, если нет - 0
	*b = bitint(testData) // меняем значение по ссылке
}

func (b *bitint) setBit(pos uint) {
	*b |= (1 << pos) // операция "или" при сравнении со вторым числом, у которого 1 бит сдвинут на указанное количество шагов.
}
