package list


import "fmt"


type List struct {
	len       int64
	firstNode *node
}


func NewList() (l *List) {
	newList := &List{}
	return newList
}


func (l *List) Len() (len int64) {
	return l.len
}


func (l *List) Add(value int64) (index int64) {
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
}


func (l *List) RemoveByIndex(index int64) {
	if index < 0 || index >= l.len {
		return
	} 

	if index == 0 {
		l.firstNode = l.firstNode.next
	} else {
		current := l.firstNode
		for i := int64(0); i < index-1; i++ {
			current = current.next
		}
		current.next = current.next.next
	}

	l.len--
}


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
}


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
}


func (l *List) GetByIndex(index int64) (value int64, ok bool) {
	if index < 0 || index >= l.len {
		return 0, false 
	}

	current := l.firstNode
	for i := int64(0); i < index; i++ {
		current = current.next
	}

	return current.value, true
}


func (l *List) GetByValue(value int64) (index int64, ok bool) {
	current := l.firstNode
	index = int64(0)

	for current != nil {
		if current.value == value {
			return index, true
		} 
		current = current.next
		index++
	}

	return 0, false
}


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
}


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
}


func (l *List) Clear() {
	l.firstNode = nil
	l.len = 0
}


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
