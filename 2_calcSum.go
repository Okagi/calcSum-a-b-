package main

import "fmt"

type Number struct {
	value int
}

func (n Number) calcSum(a, b int) Number {
	sum := a + b
	return Number{value: sum}
}

func main() {

	var a, b int
	fmt.Println("Введите значения a и b:")
	fmt.Print("a = ")
	fmt.Scanln(&a)
	fmt.Print("b = ")
	fmt.Scanln(&b)

	num := Number{}
	result := num.calcSum(a, b)

	fmt.Println(result.value)
}
