package main

import (
	"fmt"
)

// В этом примере мы создали адаптер Adapter, который реализует интерфейс NewSystem, используя структуру OldSystem.
// Теперь мы можем использовать Adapter для вызова NewOperation, и он будет делегировать вызов OldOperation старой системе,
// обеспечивая таким образом совместимость между двумя интерфейсами.

// Интерфейс для новой системы
type NewSystem interface {
	NewOperation()
}

// Структура представляющая старую систему
type OldSystem struct {
}

// Метод для старой системы
func (os *OldSystem) OldOperation() {
	fmt.Println("Выполняется старая операция")
}

// Адаптер для старой системы, чтобы она соответствовала интерфейсу NewSystem
type Adapter struct {
	oldSystem *OldSystem
}

// Реализация метода NewOperation для адаптера
func (a *Adapter) NewOperation() {
	fmt.Println("Адаптер вызывает старую операцию")
	a.oldSystem.OldOperation()
}

func main() {
	oldSystem := &OldSystem{}
	adapter := &Adapter{oldSystem}

	fmt.Println("Выполнение операции через новую систему:")
	adapter.NewOperation()
}
