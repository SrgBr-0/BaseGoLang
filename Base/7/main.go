//7. Создайте программу, вычисляющую факториал числа,
//введенного пользователем.

package main

import (
	"fmt"
	"math/big"
)

/*
func factorial(x int) int {
	num := 1
	for i := 1; i <= x; i++ {
		num *= i
	}
	return num
}
*/

func main() {
	var n int64
	fmt.Println("Введите число:")
	fmt.Scan(&n)

	x := new(big.Int)
	x.MulRange(1, n)
	fmt.Printf("Значение факториала для числа %d = %d", n, x)
}
