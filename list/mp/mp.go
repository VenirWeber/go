package mp

import (
	"errors"
	"fmt"
	"sort"
)

type Map struct {
	key int64
	mp  map[int64]int64
}

// ErrIndexOutOfRange - ошибка, которая возвращается, если индекс выходит за границы карты.
var ErrKeyOutOfRange = errors.New("key out of range")

// ErrValueOutOfRange - ошибка, которая возвращается, если значение не содержится в карте.
var ErrValueOutOfRange = errors.New("value out of range")

// ErrValueOutOfRange - ошибка, которая возвращается, если карта пуста.
var ErrMapEmpty = errors.New("map is empty")

// Функция NewMap() создает новую структуру Map
func NewMap() *Map {
	newMap := &Map{
		mp: make(map[int64]int64),
	}
	fmt.Println("Новая карта создана")
	return newMap
}

// Len возвращает длину Map
func (mp *Map) Len() int64 {
	length := int64(len(mp.mp))
	fmt.Printf("Длина карты: %d\n", length)
	return length
}

// Add добавляет элемент в Map и возвращает его key
func (mp *Map) Add(value int64) int64 {
	mp.key++
	key := mp.key
	mp.mp[key] = value
	fmt.Printf("Элемент %v добавлен в карту с ключом %d\n", value, key)
	return key
}

// RemoveByKey удаляет элемент из Map по ключу
func (mp *Map) RemoveByKey(key int64) error {
	if len(mp.mp) == 0 {
		fmt.Printf("Error: %v\n", ErrMapEmpty)
		return ErrMapEmpty
	}
	if _, exists := mp.mp[key]; !exists {
		fmt.Printf("Error: %v\n", ErrKeyOutOfRange)
		return ErrKeyOutOfRange
	}
	delete(mp.mp, key)
	fmt.Printf("Элемент с ключом %d удален\n", key)
	return nil
}

// RemoveByValue удаляет элемент из Map по значению
func (mp *Map) RemoveByValue(value int64) error {
	if len(mp.mp) == 0 {
		fmt.Printf("Error: %v\n", ErrMapEmpty)
		return ErrMapEmpty
	}
	var keyToDelete int64
	found := false
	for key, v := range mp.mp {
		if v == value {
			keyToDelete = key
			found = true
			break
		}
	}
	if !found {
		fmt.Printf("Error: %v\n", ErrValueOutOfRange)
		return ErrValueOutOfRange
	}
	delete(mp.mp, keyToDelete)
	fmt.Printf("Элемент со значением %d удален с ключом %d\n", value, keyToDelete)
	return nil
}

// RemoveAllByValue удаляет все элементы из Map по значению
func (mp *Map) RemoveAllByValue(value int64) error {
	if len(mp.mp) == 0 {
		fmt.Printf("Error: %v\n", ErrMapEmpty)
		return ErrMapEmpty
	}
	var deletedKeys []int64
	for key, v := range mp.mp {
		if v == value {
			deletedKeys = append(deletedKeys, key)
			delete(mp.mp, key)
		}
	}
	if len(deletedKeys) > 0 {
		fmt.Printf("Элементы со значением %d и ключами %v были удалены\n", value, deletedKeys)
		return nil
	}
	fmt.Printf("Error: %v\n", ErrValueOutOfRange)
	return ErrValueOutOfRange
}

// GetByKey возвращает значение элемента по ключу.
//
// Если элемента с таким ключом нет, то возвращается 0 и false.
func (mp *Map) GetByKey(key int64) (int64, error) {
	if len(mp.mp) == 0 {
		fmt.Printf("Error: %v\n", ErrMapEmpty)
		return 0, ErrMapEmpty
	}
	value, ok := mp.mp[key]
	if !ok {
		fmt.Printf("Error: %v\n", ErrKeyOutOfRange)
		return 0, ErrKeyOutOfRange
	}
	return value, nil
}

// GetByValue возвращает ключ первого найденного элемента по значению.
//
// Если элемента с таким значением нет, то возвращается 0 и false.
func (mp *Map) GetByValue(value int64) (int64, error) {
	if len(mp.mp) == 0 {
		fmt.Printf("Error: %v\n", ErrMapEmpty)
		return 0, ErrMapEmpty
	}
	for key, v := range mp.mp {
		if v == value {
			fmt.Printf("Элемент с ключом %d и значением %d найден\n", key, value)
			return key, nil
		}
	}
	fmt.Printf("Error: %v\n", ErrValueOutOfRange)
	return 0, ErrValueOutOfRange
}

// GetAllByValue возвращает индексы всех найденных элементов по значению
//
// Если элементов с таким значением нет, то возвращается nil и false.
func (mp *Map) GetAllByValue(value int64) ([]int64, error) {
	if len(mp.mp) == 0 {
		fmt.Printf("Error: %v\n", ErrMapEmpty)
		return nil, ErrMapEmpty
	}
	var keys []int64
	for key, v := range mp.mp {
		if v == value {
			keys = append(keys, key)
		}
	}
	if len(keys) > 0 {
		fmt.Printf("Элементы со значением %d найдены с ключами %v\n", value, keys)
		return keys, nil
	}
	fmt.Printf("Error: %v\n", ErrValueOutOfRange)
	return nil, ErrValueOutOfRange
}

// GetAll возвращает все элементы списка
//
// Если список пуст, то возвращается nil и false.
func (mp *Map) GetAll() ([]int64, error) {
	if len(mp.mp) == 0 {
		fmt.Printf("Error: %v\n", ErrMapEmpty)
		return nil, ErrMapEmpty
	}
	var values []int64
	for _, v := range mp.mp {
		values = append(values, v)
	}
	fmt.Println("Все элементы были успешно получены")
	return values, nil
}

// Clear очищает список
func (mp *Map) Clear() {
	mp.mp = make(map[int64]int64)
	fmt.Println("Карта была успешно очищена")
}

// Print выводит список в консоль
func (mp *Map) Print() {
	var keys []int64

	if len(mp.mp) == 0 {
		fmt.Println("Карта пуста")
		return
	}

	for key := range mp.mp {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	fmt.Println("Содержимое Map:")

	for _, key := range keys {
		value := mp.mp[key]
		fmt.Printf("Ключ: %d, Значение: %d\n", key, value)
	}
}
