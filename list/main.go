package main

import (
	"list/storage/list"
	"list/storage/mp"
)

func main() {
	l := list.NewList()
	l.Add(11)
	l.Add(12)
	l.Add(13)
	l.Add(14)
	l.RemoveByIndex(3)

	m := mp.NewMap()
	m.Add(100)
	m.Add(200)
	m.Add(300)

	m.Print()

}
