package storage

import (
	"errors"
	"fmt"
)

type Storage struct {
	data      []interface{}
	valueType interface{}
}

var ErrMismatchType = errors.New("mismatched type: the type of the provided value does not match the type of items already in the storage")

func NewStorage() *Storage {
	return &Storage{}
}

func (s *Storage) Len() int64 {
	return int64(len(s.data))
}

func (s *Storage) Add(value interface{}) (int64, error) {
	if len(s.data) == 0 {
		s.valueType = value
	}

	if s.valueType != nil && s.valueType != value {
		return 0, ErrMismatchType
	}

	s.data = append(s.data, value)
	return int64(len(s.data) - 1), nil
}

func (s *Storage) RemoveByIndex(id int64) {
	if id < 0 || int(id) >= len(s.data) {
		return
	}
	s.data = append(s.data[:id], s.data[id+1:]...)
}

func (s *Storage) RemoveByValue(value interface{}) {
	for i, v := range s.data {
		if v == value {
			s.RemoveByIndex(int64(i))
			break
		}
	}
}

func (s *Storage) RemoveAllByValue(value interface{}) {
	var newData []interface{}
	for _, v := range s.data {
		if v != value {
			newData = append(newData, v)
		}
	}
	s.data = newData
}

func (s *Storage) GetByIndex(id int64) (interface{}, bool) {
	if id < 0 || int(id) >= len(s.data) {
		return nil, false
	}
	return s.data[id], true
}

func (s *Storage) GetByValue(value interface{}) (int64, bool) {
	for i, v := range s.data {
		if v == value {
			return int64(i), true
		}
	}
	return 0, false
}

func (s *Storage) GetAllByValue(value interface{}) ([]int64, bool) {
	var indices []int64
	for i, v := range s.data {
		if v == value {
			indices = append(indices, int64(i))
		}
	}
	if len(indices) == 0 {
		return nil, false
	}
	return indices, true
}

func (s *Storage) GetAll() ([]interface{}, bool) {
	if len(s.data) == 0 {
		return nil, false
	}
	return s.data, true
}

func (s *Storage) Clear() {
	s.data = nil
}

func (s *Storage) Print() {
	fmt.Println(s.data)
}
