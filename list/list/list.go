package list

import "fmt"

type List struct {
	len       int64
	firstNode *node
}

// Создаётся новый список и возвращает указатель
func NewList() (l *List) {
	newList := &List{}
	return newList
}

// Возращает длину списка, которая хранится в поле len структуры List
func (l *List) Len() (len int64) {
	return l.len
}

// Добавляет элемент в список и возвращает его индекс
func (l *List) Add(value int64) (index int64) {
	newNode := &node{
		value: value,
		next:  nil,
	}
	/* Создаем новый узел с заданным значением value
	и добавляем его в конец списка. */

	if l.firstNode == nil {
		l.firstNode = newNode
	} else {
		current := l.firstNode
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
	/*Если список пустой, устанавливаем новый узел как первый узел.
	В противном случае, мы перебираем список, пока не найдем последний узел,
	и добавляем новый узел в конец списка.*/

	l.len++
	return l.len - 1
	/* Возвращает индекс добавленного элемента,
	который равен текущей длине списка - 1 (т.к. индексация начинается с 0).*/
}

// Удаляет элемент из списка по индексу
func (l *List) RemoveByIndex(index int64) {
	if index < 0 || index >= l.len {
		return
	} // Проверяем, что индекс находится в допустимых пределах

	if index == 0 {
		// Если удаляемый элемент - первый в списке, переназначаем первый узел
		l.firstNode = l.firstNode.next
	} else {
		/* В противном случае, находим узел перед удаляемым узлом
		   и переназначаем указатель на следующий узел*/
		current := l.firstNode
		for i := int64(0); i < index-1; i++ {
			current = current.next
		}
		current.next = current.next.next
	}

	l.len-- /* Удаляем элемент из списка, обновляя указатели
	узлов таким образом, чтобы исключить удаляемый элемент.*/
}

// Удаление элемента из списка по значению
func (l *List) RemoveByValue(value int64) {
	if l.len == 0 {
		return
	} // Обработка случая, когда список пуст

	if l.firstNode.value == value {
		l.firstNode = l.firstNode.next
		l.len--
		return
	} // Если удаляемый элемент находится в начале списка

	current := l.firstNode
	for current.next != nil && current.next.value != value {
		current = current.next
	} // Поиск элемента для удаления

	if current.next == nil {
		return
	} // Если элемент не найден, выходим

	// Удаляем элемент, переназначая указатель на следующий узел
	current.next = current.next.next
	l.len--
}

// Удаляет все элементы из списка по значению
func (l *List) RemoveAllByValue(value int64) {
	if l.len == 0 {
		return
	} // Обработка случая, когда список пуст

	for l.firstNode != nil && l.firstNode.value == value {
		l.firstNode = l.firstNode.next
		l.len--
	} // Удаляем нужные элементы в начале списка если такие есть

	current := l.firstNode
	for current != nil && current.next != nil {
		if current.next.value == value {
			current.next = current.next.next
			l.len--
		} else {
			current = current.next
		}
	} // Поиск и удаление остальных элементов
}

// Возвращает значение элемента по индексу.
func (l *List) GetByIndex(index int64) (value int64, ok bool) {
	if index < 0 || index >= l.len {
		return 0, false // Возвращаем 0 и false, если индекс недопустим
	}

	current := l.firstNode
	for i := int64(0); i < index; i++ {
		current = current.next
	}

	return current.value, true
	// Возвращаем значение и true, если элемент найден
}

// Возвращает индекс первого найденного элемента по значению.
func (l *List) GetByValue(value int64) (index int64, ok bool) {
	current := l.firstNode
	index = int64(0)

	for current != nil {
		if current.value == value {
			return index, true
		} // Возвращаем значение и true, если элемент найден
		current = current.next
		index++
	}

	return 0, false
	// Возвращаем 0 и false, если элемент с указанным значением не найден
}

// Возвращает индексы всех найденных элементов по значению
func (l *List) GetAllByValue(value int64) (ids []int64, ok bool) {
	current := l.firstNode

	for index := int64(0); current != nil; index++ {
		if current.value == value {
			ids = append(ids, index)
		}
		current = current.next
	}

	if len(ids) > 0 {
		return ids, true
	}

	return nil, false
	/* Возвращаем nil и false,
	если элементы с указанным значением не найдены*/
}

// Возвращает все элементы списка
func (l *List) GetAll() (values []int64, ok bool) {
	if l.len == 0 {
		return nil, false // Возвращаем nil и false, если список пуст
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
}

// Очищает список
func (l *List) Clear() {
	l.firstNode = nil
	l.len = 0
}

// Выводит список в консоль
func (l *List) Print() {
	current := l.firstNode

	fmt.Print("Список: \n")
	index := 0

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
