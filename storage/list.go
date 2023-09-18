package list

type List struct {
	len       int64
	firstNode *node
}

func NewList() *List {
	newList := &List{}
	return newList
} // создаётся новый список и возвращает указатель

func (l *List) Len() (len int64) {
	return l.len
} // возращает длину списка, которая хранится в поле len структуры List

func (l *List) Add(value int64) (id int64) {
	newNode := &node{
		value: value,
		next:  nil,
	} /* создаем новый узел с заданным значением value
	и добавляем его в конец списка. */

	if l.firstNode == nil {
		l.firstNode = newNode
	} else {
		current := l.firstNode
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	} /*Если список пустой, устанавливаем новый узел как первый узел.
	В противном случае, мы перебираем список, пока не найдем последний узел,
	и добавляем новый узел в конец списка.*/

	l.len++
	return l.len - 1 /* Возвращает индекс добавленного элемента,
	который равен текущей длине списка - 1 (т.к. индексация начинается с 0).*/
} // добавляет элемент в список и возвращает его индекс

func (l *List) RemoveByIndex(id int64) {
	if id < 0 || id >= l.len {
		return
	} // Проверяем, что индекс находится в допустимых пределах

	if id == 0 {
		// Если удаляемый элемент - первый в списке, переназначаем первый узел
		l.firstNode = l.firstNode.next
	} else {
		/* В противном случае, находим узел перед удаляемым узлом
		   и переназначаем указатель на следующий узел*/
		current := l.firstNode
		for i := int64(0); i < id-1; i++ {
			current = current.next
		}
		current.next = current.next.next
	}

	l.len-- /* удаляем элемент из списка, обновляя указатели
	узлов таким образом, чтобы исключить удаляемый элемент.*/
} // удаляет элемент из списка по индексу

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
} // удаляение элемента из списка по значению

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
} // Удаляет все элементы из списка по значению

func (l *List) GetByIndex(id int64) (value int64, ok bool) {
	if id < 0 || id >= l.len {
		return 0, false // Возвращаем 0 и false, если индекс недопустим
	}

	current := l.firstNode
	for i := int64(0); i < id; i++ {
		current = current.next
	}

	return current.value, true
	// Возвращаем значение и true, если элемент найден
} // Возвращает значение элемента по индексу.

func (l *List) GetByValue(value int64) (id int64, ok bool) {
	current := l.firstNode
	id = int64(0)

	for current != nil {
		if current.value == value {
			return id, true
		} // Возвращаем значение и true, если элемент найден
		current = current.next
		id++
	}

	return 0, false
	// Возвращаем 0 и false, если элемент с указанным значением не найден
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
	/* Возвращаем nil и false,
	если элементы с указанным значением не найдены*/
} // Возвращает индексы всех найденных элементов по значению

// GetAll возвращает все элементы списка
//
// Если список пуст, то возвращается nil и false.
func (l *List) GetAll() (values []int64, ok bool) {
	return
}

// Clear очищает список
func (l *List) Clear() {

}

// Print выводит список в консоль
func (l *List) Print() {

}
