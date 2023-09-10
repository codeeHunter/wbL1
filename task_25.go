package main

import (
	"fmt"
	"strings"
)

func areAllCharactersUnique(s string) bool {
	// Приводим строку к нижнему регистру, чтобы сделать проверку регистронезависимой
	s = strings.ToLower(s)

	// Создаем карту для отслеживания уникальных символов
	charMap := make(map[rune]bool)

	for _, char := range s {
		// Если символ уже встречался, возвращаем false
		if charMap[char] {
			return false
		}
		// В противном случае помечаем его как встреченный
		charMap[char] = true
	}

	// Если мы добрались до этой точки, значит, все символы уникальны
	return true
}

func main() {
	// Примеры тестов
	fmt.Println(areAllCharactersUnique("abcd"))      // true
	fmt.Println(areAllCharactersUnique("abCdefAaf")) // false
	fmt.Println(areAllCharactersUnique("aabcd"))     // false
}
