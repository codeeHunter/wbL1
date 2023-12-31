// Данная программа выведет "0".
//
// Причина заключается в том, что внутри блока if создается новая переменная n с помощью оператора :=, и эта переменная имеет область видимости только внутри этого блока.
// Таким образом, переменная n, объявленная внутри блока if, перекрывает переменную n, объявленную в функции main.
//
// Когда программа пытается выполнить fmt.Println(n) за пределами блока if, она обращается к переменной n из функции main, которая остается неинициализированной и по умолчанию имеет значение 0.
// Внутренняя переменная n из блока if завершает свое существование после выхода из блока if, и она не имеет никакого влияния на внешнюю переменную n.
//
// Таким образом, программа выводит "0", потому что переменная n из функции main не изменяется внутри блока if.
package main

import (
	"fmt"
)

func main() {
	n := 0
	if true {
		n := 1
		n++
	}
	fmt.Println(n)
}
