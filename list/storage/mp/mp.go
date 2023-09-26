package mp

import (
	"fmt"
	"strconv"
)

type Map struct {
	data map[string]interface{}
}

// создание Map с ключами строкового типа и значениями интерфейсного типа
func NewMap() *Map {
	return &Map{
		data: make(map[string]interface{}),
	}
}

// Len возвращает длину Map
func (mp *Map) Len() int64 {
	return int64(len(mp.data))
}

// Add добавляет элемент в Map и возвращает его ключ
func (mp *Map) Add(value string) int64 {
	key := int64(len(mp.data))
	mp.data[fmt.Sprint(key)] = value
	return key
} 

// RemoveByIndex удаляет элемент из Map по ключу
func (mp *Map) RemoveByKey(key int64) {
	delete(mp.data, fmt.Sprint(key))
}

// RemoveByValue удаляет элемент из Map по значению
func (mp *Map) RemoveByValue(value string) {
	for key, val := range mp.data {
		if val == value {
			delete(mp.data, key)
			return 
		}
	}
}

// RemoveAllByValue удаляет все элементы из Map по значению
func (mp *Map) RemoveAllByValue(value string) {
	keysToRemove := make([]string, 0)

	for key, val := range mp.data {
		if val == value {
			keysToRemove = append(keysToRemove, key)
		}
	}

	for _, key := range keysToRemove {
		delete(mp.data, key)
	}
} 

// GetByKey возвращает значение элемента по ключу.
//
// Если элемента с таким ключом нет, то возвращается 0 и false.
func (mp *Map) GetByKey(key int64) (value string, ok bool) {
	val, found := mp.data[fmt.Sprint(key)]
	if !found {
		return "0", false
	}
	value, ok = val.(string)
	return value, ok
} 

// GetByValue возвращает ключ первого найденного элемента по значению.
//
// Если элемента с таким значением нет, то возвращается 0 и false.
func (mp *Map) GetByValue(value string) (key int64, ok bool) {
	for k, v := range mp.data {
		if vStr, isString := v.(string); isString && vStr == value {
			key, ok = parseKey(k)
			return
		}
	}
	return 0, false
} 

// Функция parseKey используется для преобразования строки в int64.
func parseKey(keyStr string) (int64, bool) {
	key, err := strconv.ParseInt(keyStr, 10, 64)
	if err != nil {
		return 0, false
	}
	return key, true
}

// GetAllByValue возвращает индексы всех найденных элементов по значению
//
// Если элементов с таким значением нет, то возвращается nil и false.
func (mp *Map) GetAllByValue(value string) (ids []int64, ok bool) {
	for k, v := range mp.data {
		if vStr, isString := v.(string); isString && vStr == value {
			key, parseOk := parseKey(k)
			if parseOk {
				ids = append(ids, key)
			}
		}
	}
	return ids, len(ids) > 0
}

// GetAll возвращает все элементы списка
//
// Если список пуст, то возвращается nil и false.
func (mp *Map) GetAll() (values []string, ok bool) {
	if len(mp.data) == 0 {
		return nil, false
	}

	for _, v := range mp.data {
		if vStr, isString := v.(string); isString {
			values = append(values, vStr)
		}
	}

	return values, true
}

// Clear очищает Map
func (mp *Map) Clear() {
	mp.data = make(map[string]interface{})
}

// Print выводит Map в консоль
func (mp *Map) Print() {
	fmt.Print("Вывод Map: \n")
	for key, value := range mp.data {
		fmt.Printf("Ключ: %s, Значение: %v\n", key, value)
	}
}
