package main

import (
	"fmt"
	"strings"
)

func reverseWordsInString(input string) string {
	// Разбиваем строку на слова
	words := strings.Fields(input)

	// Создаем пустую строку, в которую будем добавлять перевернутые слова
	reversedWords := make([]string, len(words))

	// Итерируемся по словам в обратном порядке и добавляем их в результат
	for i, j := 0, len(words)-1; j >= 0; i, j = i+1, j-1 {
		reversedWords[i] = words[j]
	}

	// Собираем перевернутые слова обратно в строку
	result := strings.Join(reversedWords, " ")

	return result
}

func main() {
	input := "sun dog snow"
	reversed := reverseWordsInString(input)
	fmt.Println("Исходная строка:", input)
	fmt.Println("Строка с перевернутыми словами:", reversed)
}
