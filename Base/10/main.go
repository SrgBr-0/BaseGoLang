//10. Напишите программу, которая принимает на вход список имен
//и сортирует их по алфавиту. Для имен на Русском языке!

package main

import (
	"bufio"
	"fmt"
	"golang.org/x/text/collate"
	"golang.org/x/text/language"
	"os"
	"strings"
)

func main() {
	fmt.Println("Введите список имен:")
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	words := strings.Fields(text)
	c := collate.New(language.Russian, collate.IgnoreCase)
	c.SortStrings(words)
	fmt.Println("Отсортированные имена по алфавиту: ", words)
}
