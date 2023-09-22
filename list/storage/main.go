package main

import list "list/storage/list"

func main() {
	l := list.NewList()
	l.Add(11)
	l.Add(12)
	l.Add(13)
	l.Add(14)
	l.RemoveByIndex(3)
	l.Print()
}
