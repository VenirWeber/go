package list

import "fmt"

type List struct {
	len       int64
	firstNode *node
}

func NewList() (l *List) {
	newList := &List{}
	return newList
} // Создаётся новый список и возвращает указатель

func (l *List) Len() (len int64) {
	return l.len
} // Возращает длину списка, которая хранится в поле len структуры List

func (l *List) Add(value int64) (id int64) {
	newNode := &node{
		value: value,
		next:  nil,
	}

	if l.firstNode == nil {
		l.firstNode = newNode
	} else {
		current := l.firstNode
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}

	l.len++
	return l.len - 1
} // Добавляет элемент в список и возвращает его индекс

func (l *List) RemoveByIndex(id int64) {
	if id < 0 || id >= l.len {
		return
	}

	if id == 0 {
		l.firstNode = l.firstNode.next
	} else {
		current := l.firstNode
		for i := int64(0); i < id-1; i++ {
			current = current.next
		}
		current.next = current.next.next
	}

	l.len--
} // Удаляет элемент из списка по индексу

func (l *List) RemoveByValue(value int64) {
	if l.len == 0 {
		return
	}

	if l.firstNode.value == value {
		l.firstNode = l.firstNode.next
		l.len--
		return
	}

	current := l.firstNode
	for current.next != nil && current.next.value != value {
		current = current.next
	}

	if current.next == nil {
		return
	}

	current.next = current.next.next
	l.len--
} // Удаление элемента из списка по значению

func (l *List) RemoveAllByValue(value int64) {
	if l.len == 0 {
		return
	}

	for l.firstNode != nil && l.firstNode.value == value {
		l.firstNode = l.firstNode.next
		l.len--
	}

	current := l.firstNode
	for current != nil && current.next != nil {
		if current.next.value == value {
			current.next = current.next.next
			l.len--
		} else {
			current = current.next
		}
	}
} // Удаляет все элементы из списка по значению

func (l *List) GetByIndex(id int64) (value int64, ok bool) {
	if id < 0 || id >= l.len {
		return 0, false
	}

	current := l.firstNode
	for i := int64(0); i < id; i++ {
		current = current.next
	}

	return current.value, true
} // Возвращает значение элемента по индексу.

func (l *List) GetByValue(value int64) (id int64, ok bool) {
	current := l.firstNode
	id = int64(0)

	for current != nil {
		if current.value == value {
			return id, true
		}
		current = current.next
		id++
	}

	return 0, false
} // Возвращает индекс первого найденного элемента по значению.

func (l *List) GetAllByValue(value int64) (ids []int64, ok bool) {
	current := l.firstNode

	for id := int64(0); current != nil; id++ {
		if current.value == value {
			ids = append(ids, id)
		}
		current = current.next
	}

	if len(ids) > 0 {
		return ids, true
	}

	return nil, false
} // Возвращает индексы всех найденных элементов по значению

func (l *List) GetAll() (values []int64, ok bool) {
	if l.len == 0 {
		return nil, false
	}

	values = make([]int64, l.len)
	current := l.firstNode
	index := 0

	for current != nil {
		values[index] = current.value
		current = current.next
		index++
	}

	return values, true
} // Возвращает все элементы списка

func (l *List) Clear() {
	l.firstNode = nil
	l.len = 0
} // Очищает список

func (l *List) Print() {
	current := l.firstNode

	fmt.Print("Список: \n")
	index := 0

	for current != nil {
		fmt.Print("id = ")
		fmt.Printf("%v  ", index)
		fmt.Print("значение = ")
		fmt.Printf("%v  ", current.value)
		fmt.Print("\n")
		index++
		current = current.next
	}

	fmt.Println()
} // Выводит список в консоль
