package storage

import (
	"errors"
	"fmt"
)

type Storage struct {
	data      []interface{}
	valueType interface{}
}

// ErrIndexOutOfRange - ошибка, которая возвращается, если индекс выходит за границы контейнера.
var ErrIndexOutOfRange = errors.New("key out of range")

// ErrValueOutOfRange - ошибка, которая возвращается, если значение не содержится в контейнере.
var ErrValueOutOfRange = errors.New("value out of range")

// ErrValueOutOfRange - ошибка, которая возвращается, если контейнер пуст.
var ErrStorageEmpty = errors.New("storage is empty")

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
		fmt.Printf("Error: %v\n", ErrMismatchType)
		return 0, ErrMismatchType
	}

	s.data = append(s.data, value)
	return int64(len(s.data) - 1), nil
}

func (s *Storage) RemoveByIndex(id int64) error {
	if len(s.data) == 0 {
		fmt.Printf("Error: %v\n", ErrStorageEmpty)
		return ErrStorageEmpty
	}

	if id < 0 || id >= int64(len(s.data)) {
		fmt.Printf("Error: %v\n", ErrIndexOutOfRange)
		return ErrIndexOutOfRange
	}

	if id < int64(len(s.data)-1) {
		s.data = append(s.data[:id], s.data[id+1:]...)
	} else {
		s.data = s.data[:id]
	}

	fmt.Printf("Элемент %v с индексом %d был удален\n", s.data[id], id)
	return nil
}

func (s *Storage) RemoveByValue(value interface{}) error {
	if len(s.data) == 0 {
		fmt.Println("Error:", ErrStorageEmpty)
		return ErrStorageEmpty
	}

	var removed bool
	for i, v := range s.data {
		if v == value {
			if err := s.RemoveByIndex(int64(i)); err != nil {
				fmt.Println("Error:", err)
				return err
			}
			fmt.Printf("Элемент %v с индексом %d был удален\n", value, i)
			removed = true
		}
	}

	if !removed {
		fmt.Println("Error:", ErrValueOutOfRange)
		return ErrValueOutOfRange
	}

	return nil
}

func (s *Storage) RemoveAllByValue(value interface{}) {
	var newData []interface{}
	removedIndices := make([]int64, 0)
	for i, v := range s.data {
		if v != value {
			newData = append(newData, v)
		} else {
			removedIndices = append(removedIndices, int64(i))
		}
	}
	s.data = newData
	if len(removedIndices) > 0 {
		fmt.Printf("Элемент %v с индексом %d был удален\n", value, removedIndices)
	}
}

func (s *Storage) GetByIndex(id int64) (interface{}, error) {
	if id < 0 || int(id) >= len(s.data) {
		fmt.Printf("Error: %v\n", ErrIndexOutOfRange)
		return nil, ErrIndexOutOfRange
	}
	return s.data[id], nil
}

func (s *Storage) GetByValue(value interface{}) (int64, error) {
	for i, v := range s.data {
		if v == value {
			fmt.Printf("Элемент с значением %v найден с индексом %d\n", value, i)
			return int64(i), nil
		}
	}
	fmt.Printf("Ошибка: %v не найден\n", ErrValueOutOfRange)
	return 0, ErrValueOutOfRange
}

func (s *Storage) GetAllByValue(value interface{}) ([]int64, error) {
	var indices []int64
	for i, v := range s.data {
		if v == value {
			indices = append(indices, int64(i))
			fmt.Printf("Элемент со значением %v найден по индексу %d\n", value, i)
		}
	}
	if len(indices) == 0 {
		fmt.Printf("Error: %v\n", ErrValueOutOfRange)
		return nil, ErrValueOutOfRange
	}
	return indices, nil
}

func (s *Storage) GetAll() ([]interface{}, error) {
	if len(s.data) == 0 {
		fmt.Printf("Error: %v\n", ErrStorageEmpty)
		return nil, ErrStorageEmpty
	}
	return s.data, nil
}

func (s *Storage) Clear() {
	s.data = nil
	fmt.Println("Хранилище было успешно очищено")
}

func (s *Storage) Print() {
	fmt.Print("Содержимое Storage: ")
	for i, item := range s.data {
		fmt.Printf("Индекс: %d, Элемент: %v;", i, item)
	}
	fmt.Println()
}
