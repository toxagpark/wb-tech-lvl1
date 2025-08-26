package main

import "fmt"

func setBit(a int64, bit, i int) int64 {
	if bit == 1 { // Нам нужно поменять в 1. Перемещаем 1 на нужную позицию.
		// Нужная позиция в a станет 1, так как |
		return a | (1 << i)
	} else {
		//почти так же, только с помощью ^ меняем нашу позицию в маске на 0 а все остальное становится 1
		// все остается так же, как и было кроме i ==> i - 0
		mask := int64(1) << i
		return a & ^mask
	}
}

func main() {
	var a int64
	a = 5
	bit := 0
	i := 0
	fmt.Println(setBit(a, bit, i))
}
