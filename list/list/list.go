package list

import (
	"errors"
	"fmt"
)

// ErrIndexOutOfRange - ошибка, которая возвращается, если индекс выходит за границы списка.
var ErrIndexOutOfRange = errors.New("index out of range")

// ErrValueOutOfRange - ошибка, которая возвращается, если значение не содержится в списке.
var ErrValueOutOfRange = errors.New("value out of range")

// ErrValueOutOfRange - ошибка, которая возвращается, если список пуст.
var ErrListEmpty = errors.New("list is empty")

type List struct {
	len       int64
	firstNode *node
}

// Создаётся новый список и возвращает указатель
func NewList() *List {
	fmt.Println("Создан новый список")
	return &List{}
}

// Возращает длину списка, которая хранится в поле len структуры List
func (l *List) Len() int64 {
	fmt.Printf("Длина списка: %d\n", l.len)
	return l.len
}

// Добавляет элемент в список и возвращает его индекс
func (l *List) Add(value int64) int64 {
	newNode := &node{
		value: value,
		next:  nil,
	}

	if l.firstNode == nil {
		l.firstNode = newNode
		fmt.Printf("Элемент %v добавлен в список с индексом 0\n", value)
	} else {
		current := l.firstNode
		index := int64(0)
		for current.next != nil {
			current = current.next
			index++
		}
		current.next = newNode
		fmt.Printf("Элемент %v добавлен в список с индексом %d\n", value, index+1)
	}

	l.len++
	return l.len - 1
}

// Удаляет элемент из списка по индексу
func (l *List) RemoveByIndex(index int64) error {
	if l.len == 0 {
		fmt.Printf("Error: %v\n", ErrListEmpty)
		return ErrListEmpty
	}

	if index < 0 || index >= l.len {
		fmt.Printf("Error: %v\n", ErrIndexOutOfRange)
		return ErrIndexOutOfRange
	}

	var deletedValue int64

	if index == 0 {
		deletedValue = l.firstNode.value
		l.firstNode = l.firstNode.next
	} else {
		current := l.firstNode
		for i := int64(0); i < index-1; i++ {
			current = current.next
		}
		deletedValue = current.next.value
		current.next = current.next.next
	}

	l.len--

	fmt.Printf("Элемент %v с индексом %d был удален\n", deletedValue, index)
	return nil
}

// Удаление элемента из списка по значению
func (l *List) RemoveByValue(value int64) error {
	if l.len == 0 {
		fmt.Printf("Error: %v\n", ErrListEmpty)
		return ErrListEmpty
	}

	var deletedIndex int64

	if l.firstNode.value == value {
		deletedIndex = 0
		l.firstNode = l.firstNode.next
		l.len--
		return nil
	} else {
		current := l.firstNode
		previous := current
		index := int64(0)

		for current != nil && current.value != value {
			previous = current
			current = current.next
			index++
		}

		if current == nil {
			fmt.Printf("Error: %v\n", ErrValueOutOfRange)
			return ErrValueOutOfRange
		}

		previous.next = current.next
		l.len--
		deletedIndex = index
	}

	fmt.Printf("Элемент %v с индексом %d был удален\n", value, deletedIndex)
	return nil
}

// Удаляет все элементы из списка по значению
func (l *List) RemoveAllByValue(value int64) error {
	if l.len == 0 {
		fmt.Printf("Error: %v\n", ErrListEmpty)
		return ErrListEmpty
	}

	var deletedIndices []int64

	for l.firstNode != nil && l.firstNode.value == value {
		deletedIndices = append(deletedIndices, 0)
		l.firstNode = l.firstNode.next
		l.len--
	}

	current := l.firstNode
	index := int64(0)
	for current != nil && current.next != nil {
		if current.next.value == value {
			deletedIndices = append(deletedIndices, index+1)
			current.next = current.next.next
			l.len--
		} else {
			current = current.next
		}
		index++
	}

	if len(deletedIndices) > 0 {
		fmt.Printf("Элементы %v с индексами %v были удалены\n", value, deletedIndices)
	}

	return nil
}

// Возвращает значение элемента по индексу.
func (l *List) GetByIndex(index int64) (int64, error) {
	if l.len == 0 {
		fmt.Printf("Error: %v\n", ErrListEmpty)
		return 0, ErrListEmpty
	}

	if index < 0 || index >= l.len {
		fmt.Printf("Error: %v\n", ErrIndexOutOfRange)
		return 0, ErrIndexOutOfRange
	}

	current := l.firstNode
	for i := int64(0); i < index; i++ {
		current = current.next
	}

	fmt.Printf("Элемент %v с индексом %d найден\n", current.value, index)
	return current.value, nil
}

// Возвращает индекс первого найденного элемента по значению.
func (l *List) GetByValue(value int64) (int64, error) {
	if l.len == 0 {
		fmt.Printf("Error: %v\n", ErrListEmpty)
		return 0, ErrListEmpty
	}
	current := l.firstNode
	index := int64(0)

	for current != nil {
		if current.value == value {
			fmt.Printf("Элемент %v с индексом %d найден\n", value, index)
			return index, nil
		}
		current = current.next
		index++
	}

	fmt.Printf("Error: %v\n", ErrValueOutOfRange)
	return 0, ErrValueOutOfRange
	// Возвращаем 0 и ErrValueOutOfRange, если элемент с указанным значением не найден
}

// Возвращает индексы всех найденных элементов по значению
func (l *List) GetAllByValue(value int64) ([]int64, error) {
	if l.len == 0 {
		fmt.Printf("Error: %v\n", ErrListEmpty)
		return nil, ErrListEmpty
	}
	current := l.firstNode
	var ids []int64
	for index := int64(0); current != nil; index++ {
		if current.value == value {
			fmt.Printf("Элемент %v с индексом %d найден\n", value, index)
			ids = append(ids, index)
		}
		current = current.next
	}

	if len(ids) > 0 {
		return ids, nil
	}

	fmt.Printf("Error: %v\n", ErrValueOutOfRange)
	return nil, ErrValueOutOfRange
}

// Возвращает все элементы списка
func (l *List) GetAll() ([]int64, error) {
	if l.len == 0 {
		fmt.Printf("Error: %v\n", ErrListEmpty)
		return nil, ErrListEmpty
	}

	var values = make([]int64, l.len)
	current := l.firstNode
	index := 0

	for current != nil {
		values[index] = current.value
		current = current.next
		index++
	}

	fmt.Println("Все элементы успешно получены")
	return values, nil
}

// Очищает список
func (l *List) Clear() {
	l.firstNode = nil
	l.len = 0
	fmt.Println("Список был очищен")
}

// Выводит список в консоль
func (l *List) Print() {
	current := l.firstNode

	fmt.Print("Список: \n")
	index := 0

	if l.len == 0 {
		fmt.Println("Список пуст.")
		return
	}

	for current != nil {
		fmt.Print("index = ")
		fmt.Printf("%v  ", index)
		fmt.Print("значение = ")
		fmt.Printf("%v  ", current.value)
		fmt.Print("\n")
		index++
		current = current.next
	}

	fmt.Println()
}
