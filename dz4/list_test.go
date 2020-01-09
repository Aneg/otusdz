package dz4

import (
	"testing"
)

func TestNewList(t *testing.T) {
	list := NewList()

	if list.Len() != 0 {
		t.Error("не верная длина пустого списка")
	}
}

func TestList_PushFront(t *testing.T) {
	list := NewList()

	list.PushFront(1)
	if list.Len() != 1 || list.Last().Value() != 1 || list.First().Value() != 1 {
		t.Error("эллемент не верно добавлен очередь")
	}

	list.PushFront(2)
	if list.Len() != 2 || list.Last().Value() != 1 || list.First().Value() != 2 {
		t.Error("эллемент не верно добавлен очередь")
	}

	list.PushFront(3)
	if list.Len() != 3 || list.Last().Value() != 1 || list.First().Value() != 3 {
		t.Error("эллемент не верно добавлен очередь")
	}
}

func TestList_PushBack(t *testing.T) {
	list := NewList()

	list.PushBack(1)
	if list.Len() != 1 || list.Last().Value() != 1 || list.First().Value() != 1 {
		t.Error("эллемент не верно добавлен очередь")
	}

	list.PushBack(2)
	if list.Len() != 2 || list.Last().Value() != 2 || list.First().Value() != 1 {
		t.Error("эллемент не верно добавлен очередь")
	}

	list.PushBack(3)
	if list.Len() != 3 || list.Last().Value() != 3 || list.First().Value() != 1 {
		t.Error("эллемент не верно добавлен очередь")
	}
}

func TestList_Remove(t *testing.T) {
	list := NewList()

	for i := 0; i < 10; i++ {
		list.PushFront(i)
	}

	if err := list.Remove(list.First()); err != nil || list.First().Value() != 8 {
		t.Fatalf("эллемент неверно удален из начала списка")
	}
	if err := list.Remove(list.Last()); err != nil || list.Last().Value() != 1 {
		t.Fatalf("эллемент неверно удален из конца списка")
	}
	if err := list.Remove(list.Last().Next()); err != nil || list.Last().Next().Value() != 3 {
		t.Fatalf("эллемент неверно удален из списка")
	}
}
