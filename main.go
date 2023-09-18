package main

import (
	"fmt"
	"https://github.com/VenirWeber/listgo/blob/main/storage/list.go"
)

func main() {
	myList := &List{}
	myList.Add(42)
	myList.Add(123)
	myList.Add(7)
	myList.Add(123) // Добавляем элемент со значением 123 второй раз

	ids, ok := myList.GetAllByValue(123) // Получение индексов элементов со значением 123

	if ok {
		fmt.Println("Индексы элементов:", ids)
	} else {
		fmt.Println("Элементы с указанным значением не найдены.")
	}

	ids, ok = myList.GetAllByValue(99) // Попытка получить индексы элементов со значением 99 (не существует)

	if ok {
		fmt.Println("Индексы элементов:", ids)
	} else {
		fmt.Println("Элементы с указанным значением не найдены.")
	}
}
