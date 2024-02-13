//6. Напишите программу, которая принимает введенную пользователем
//строку и определяет, является ли она палиндромом
//(словом или фразой, которая читается одинаково
//как в прямом, так и в обратном порядке).

package main

import (
	"fmt"
	"strings"
)

func palindromeCheck(a string) bool {
	var reversePal string
	for i := len(a) - 1; i >= 0; i-- {
		reversePal += string(a[i])
	}
	if strings.ToLower(a) == strings.ToLower(reversePal) {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println("Введите строку:")
	var pol string
	fmt.Scan(&pol)
	if palindromeCheck(pol) {
		fmt.Println("Это палиндром")
	} else {
		fmt.Println("Это не палиндром")
	}
}
