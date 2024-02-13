//5. Создать программу, генерирующую для пользователя
//случайный пароль заданной длины и набора символов.

package main

import (
	"fmt"
	"math/rand"
)

func input(num []string, err error) []string {
	if err != nil {
		return num
	}
	var d string
	n, err := fmt.Scanf("%s", &d)
	if n == 1 {
		num = append(num, d)
	}
	return input(num, err)
}

func generator(num []string, lenPass int) string {
	var pass string
	for i := 0; i < lenPass; i++ {
		randomIndex := rand.Intn(lenPass)
		pick := num[randomIndex]
		pass += pick
	}
	return pass
}

func main() {
	fmt.Println("Введите набор букв и цифр для пароля:")
	x := input([]string{}, nil)

	var l int
	fmt.Println("Введите желаемую длину пароля:")
	fmt.Scan(&l)

	fmt.Println("Итоговый пароль:", generator(x, l))
}
