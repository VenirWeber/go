package list

import (
	"errors"
	"fmt"
)

var ErrIndexOutOfRange = errors.New("index out of range")
var ErrValueOutOfRange = errors.New("value not found")
var ErrListEmpty = errors.New("list is empty")

type List struct {
	len       int64
	firstNode *node
}


func NewList() *List {
	fmt.Println("Создан новый список")
	return &List{}
}


func (l *List) Len() int64 {
	fmt.Printf("Длина списка: %d\n", l.len)
	return l.len
}


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


func (l *List) RemoveByIndex(index int64) error {
	if index < 0 || index >= l.len {
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


func (l *List) RemoveByValue(value int64) error {
	if l.len == 0 {
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
			return ErrValueOutOfRange
		} 

		previous.next = current.next
		l.len--
		deletedIndex = index
	}

	fmt.Printf("Элемент %v с индексом %d был удален\n", value, deletedIndex)
	return nil
}


func (l *List) RemoveAllByValue(value int64) error {
	if l.len == 0 {
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


func (l *List) GetByIndex(index int64) (int64, error) {
	if index < 0 || index >= l.len {
		return 0, ErrIndexOutOfRange 
	}

	current := l.firstNode
	for i := int64(0); i < index; i++ {
		current = current.next
	}

	fmt.Printf("Элемент %v с индексом %d найден\n", current.value, index)
	return current.value, nil
}


func (l *List) GetByValue(value int64) (int64, error) {
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

	return 0, ErrValueOutOfRange
	
}


func (l *List) GetAllByValue(value int64) ([]int64, error) {
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

	return nil, ErrValueOutOfRange
}


func (l *List) GetAll() ([]int64, error) {
	if l.len == 0 {
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


func (l *List) Clear() {
	l.firstNode = nil
	l.len = 0
	fmt.Println("Список был очищен")
}


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
