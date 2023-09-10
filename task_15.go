package main

// Переполнение стека: Функция someFunc создает большую строку, которая затем копируется в justString.
// Если размер создаваемой строки слишком большой,
// это может привести к переполнению стека, так как по умолчанию строки в Go являются неизменяемыми и хранятся в стеке.
// Ошибки индексации: Если createHugeString возвращает строку меньшей длины, чем 100 символов,
// то произойдет ошибка индексации при попытке получить подстроку v[:100].

import (
	"strings"
)

var justString string

func someFunc() {
	var builder strings.Builder
	v := createHugeString(1 << 10)

	// Ограничиваем длину строки до 100 символов, чтобы избежать переполнения стека
	if len(v) > 100 {
		builder.WriteString(v[:100])
	} else {
		builder.WriteString(v)
	}

	justString = builder.String()
}

func main() {
	someFunc()
}

func createHugeString(size int) string {
	// Здесь должен быть код, создающий строку заданного размера
	// Например, можно использовать strings.Repeat для создания строки заданной длины
	return strings.Repeat("x", size)
}
