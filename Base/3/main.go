//3. Создайте программу, которая принимает число
//в качестве входных данных и отображает
//таблицу умножения этого числа до 10.

package main

import (
	"fmt"
)

func main() {
	var a float32
	fmt.Println("Введите число:")
	fmt.Scan(&a)
	for i := 1; i < 11; i++ {
		mul := float32(i) * a
		fmt.Printf("При умножении на %d: %f * %d = %f\n", i, a, i, mul)
	}
}
