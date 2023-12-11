package main

import (
	"fmt"
	"math/big"
)

// func calculate[T any](a, b T) {
// 	var sum, sub, div, mul T
// 	sum.Add(&a, &b)
// 	sub.Sub(&a, &b)
// 	div.Div(&a, &b)
// 	mul.Mul(&a, &b)

// 	fmt.Println("Sum: ")
// 	fmt.Println(sum.String())
// 	fmt.Println("Subdivision: ")
// 	fmt.Println(sub.String())
// 	fmt.Println("Division: ")
// 	fmt.Println(div.String())
// 	fmt.Println("Multiplication: ")
// 	fmt.Println(mul.String())
// }

// type BigType[T any] interface {
// 	Add()
// 	Sub()
// 	Div()
// 	Mul()
// }

// func (b *BigType[T]) Add

func main() {
	var a big.Int
	a.SetString("1234567890123456789012345678901234567890", 10)

	var b big.Int
	b.SetString("9876543219876543219876534321", 10)

	var sum, sub, div, mul big.Int
	sum.Add(&a, &b)
	sub.Sub(&a, &b)
	div.Div(&a, &b)
	mul.Mul(&a, &b)

	fmt.Println("Sum: ")
	fmt.Println(sum.String())
	fmt.Println("Subdivision: ")
	fmt.Println(sub.String())
	fmt.Println("Division: ")
	fmt.Println(div.String())
	fmt.Println("Multiplication: ")
	fmt.Println(mul.String())

}
