//4. Напишите программу, которая принимает на вход список чисел
// и находит в нем наибольшее и наименьшее число.

package main

import (
	"fmt"
	"slices"
)

func input(num []int, err error) []int {
	if err != nil {
		return num
	}
	var d int
	n, err := fmt.Scanf("%d", &d)
	if n == 1 {
		num = append(num, d)
	}
	return input(num, err)
}

func maxNum(num []int) int {
	return slices.Max(num)
}

func minNum(num []int) int {
	return slices.Min(num)
}
func main() {
	fmt.Println("Введите список чисел")
	mas := input([]int{}, nil)

	fmt.Println("Максимально введенное число = ", maxNum(mas))
	fmt.Println("Минимально введенное число = ", minNum(mas))
}
