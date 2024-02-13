//8. Напишите программу, которая берет предложение
//в качестве входных данных и преобразует его в свиную латынь
//(языковая игра, в которой слова в английском языке изменяются
//путем перемещения первого согласного
//или группы согласных в конец слова и добавления «ау»). Программа для русских фраз!

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func toPigText(text string) string {
	words := strings.Fields(text)
	pigWords := make([]string, len(words))

	for i, word := range words {
		pigWords[i] = toPigWord(word)
	}

	return strings.Join(pigWords, " ")
}

func toPigWord(word string) string {
	elem := "уеыаоэяиёю"
	firstElemIndex := strings.IndexAny(word, elem)

	if firstElemIndex == -1 {
		return word + "ау"
	}

	return word[firstElemIndex:] + word[:firstElemIndex] + "ау"
}

func main() {
	fmt.Println("Введите фразу:")
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Println("Свинаю латынь этой фразы:", toPigText(text))
}
