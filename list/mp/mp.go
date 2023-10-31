package mp

import (
	"fmt"
	"sort"
)

type Map struct {
	id int64
	mp map[int64]int64
}

// Функция NewMap() создает новую структуру Map
func NewMap() *Map {
	return &Map{
		mp: make(map[int64]int64),
	}
}

// Len возвращает длину Map
func (mp *Map) Len() int64 {
	return int64(len(mp.mp))
}

// Add добавляет элемент в Map и возвращает его key
func (mp *Map) Add(value int64) int64 {
	mp.id++

	key := mp.id

	mp.mp[key] = value

	return key
} /* Эта функция увеличивает key в структуре Map на один,
затем добавляет новое значение в Map, используя увеличенное значение key
как ключ, и возвращает этот новый key */

// RemoveByKey удаляет элемент из Map по ключу
func (mp *Map) RemoveByKey(key int64) {
	delete(mp.mp, key)
} /* Функция использует функцию delete(),
чтобы удалить элемент с заданным key из Map.*/

// RemoveByValue удаляет элемент из Map по значению
func (mp *Map) RemoveByValue(value int64) {
	for key, v := range mp.mp {
		if v == value {
			delete(mp.mp, key)
			break // Удаляем только одно совпадение
		}
	}
} /* Функция перебирает все элементы Map и удаляет первый элемент
с заданным значением, который она найдет.*/

// RemoveAllByValue удаляет все элементы из Map по значению
func (mp *Map) RemoveAllByValue(value int64) {
	for key, v := range mp.mp {
		if v == value {
			delete(mp.mp, key)
		}
	}
} /* Функция перебирает все элементы Map и удаляет
все элементы с заданным значением.*/

// GetByKey возвращает значение элемента по ключу.
//
// Если элемента с таким ключом нет, то возвращается 0 и false.
func (mp *Map) GetByKey(key int64) (int64, bool) {
	value, ok := mp.mp[key]

	return value, ok
} /* Функция пытается получить значение по заданному ключу key.Если элемент
с таким key существует в Map, она вернет его значение и устанавливает ok в true.
Если элемента с таким ключом нет, она вернет 0 и false*/

// GetByValue возвращает ключ первого найденного элемента по значению.
//
// Если элемента с таким значением нет, то возвращается 0 и false.
func (mp *Map) GetByValue(value int64) (int64, bool) {
	for key, v := range mp.mp {
		if v == value {
			return key, true
		}
	}

	return 0, false
} /* Функция перебирает все элементы Map и возвращает первый
найденный key элемента с заданным значением и устанавливает ok в true.
Если элемента с таким значением нет, она возвращает 0 и false.*/

// GetAllByValue возвращает индексы всех найденных элементов по значению
//
// Если элементов с таким значением нет, то возвращается nil и false.
func (mp *Map) GetAllByValue(value int64) ([]int64, bool) {
	var ids []int64

	for key, v := range mp.mp {
		if v == value {
			ids = append(ids, key)
		}
	}

	if len(ids) > 0 {
		return ids, true
	}

	return nil, false
} /* Функция перебирает все элементы Map, добавляя key всех элементов
с заданным значением в срез ids. Если были найдены элементы,
она возвращает срез ids и устанавливает ok в true. Если элементов с таким
значением не найдено, она возвращает nil и false*/

// GetAll возвращает все элементы списка
//
// Если список пуст, то возвращается nil и false.
func (mp *Map) GetAll() ([]int64, bool) {
	if len(mp.mp) == 0 {
		return nil, false
	}

	var values []int64
	for _, v := range mp.mp {
		values = append(values, v)
	}

	return values, true
} /* Функция проверяет, пуст ли Map. Если Map пуста, она возвращает
nil и false. В противном случае, она создает срез values и
добавляет в него все значения элементов карты, затем возвращает
этот срез и устанавливает ok в true*/

// Clear очищает список
func (mp *Map) Clear() {
	mp.mp = make(map[int64]int64)
}

// Print выводит список в консоль
func (mp *Map) Print() {
	var keys []int64

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
