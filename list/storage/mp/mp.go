package mp

import (
	"fmt"
	"sort"
)

type Map struct {
	id int64
	mp map[int64]int64
}


func NewMap() *Map {
	return &Map{
		mp: make(map[int64]int64),
	}
}


func (mp *Map) Len() int64 {
	return int64(len(mp.mp))
}


func (mp *Map) Add(value int64) int64 {
	mp.id++
	key := mp.id
	mp.mp[key] = value
	return key
} 


func (mp *Map) RemoveByKey(key int64) {
	delete(mp.mp, key)
} 


func (mp *Map) RemoveByValue(value int64) {
	for key, v := range mp.mp {
		if v == value {
			delete(mp.mp, key)
			break 
		}
	}
} 


func (mp *Map) RemoveAllByValue(value int64) {
	for key, v := range mp.mp {
		if v == value {
			delete(mp.mp, key)
		}
	}
} 


func (mp *Map) GetByKey(key int64) (int64, bool) {
	value, ok := mp.mp[key]
	return value, ok
} 

func (mp *Map) GetByValue(value int64) (int64, bool) {
	for key, v := range mp.mp {
		if v == value {
			return key, true
		}
	}
	return 0, false
} 


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
} 


func (mp *Map) GetAll() ([]int64, bool) {
	if len(mp.mp) == 0 {
		return nil, false
	}

	var values []int64
	for _, v := range mp.mp {
		values = append(values, v)
	}

	return values, true
} 


func (mp *Map) Clear() {
	mp.mp = make(map[int64]int64)
}


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


