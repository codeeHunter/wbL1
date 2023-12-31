// Данная программа выведет: [100 2 3 4 5].

// Причина этого результата заключается в том, как в Go работает передача срезов (slices) в функции.

// В функции someAction, мы изменяем первый элемент среза v на 100 с помощью выражения v[0] = 100.
// Таким образом, первый элемент среза a также изменится, потому что он совместно использует память с срезом v.
// Это происходит из-за того, что срезы в Go - это ссылки на массивы, и изменения элементов среза отражаются на оригинальном массиве.

// Затем мы добавляем значение b в срез v с помощью функции append(v, b). Однако это изменение не будет отражено на срезе a, так как срез v создается как копия среза a, и дальнейшие изменения в v не влияют на a. Функция append создает новый срез и добавляет элемент b к нему, но этот новый срез не возвращается из функции и, следовательно, не сохраняется в переменной.

// Поэтому после выполнения функции someAction, срез a все еще содержит [1 2 3 4 5], и изменение среза v внутри функции влияет только на его копию.
package main

import (
	"fmt"
)

func someAction(v []int8, b int8) {
	v[0] = 100
	v = append(v, b)
}

func main() {
	var a = []int8{1, 2, 3, 4, 5}
	someAction(a, 6)
	fmt.Println(a)
}
