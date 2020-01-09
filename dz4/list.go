package dz4

import (
	"errors"
)

func NewList() *List {
	return &List{len: 0}
}

type List struct {
	first *Item
	last  *Item
	len   uint
}

func (l *List) Last() *Item {
	return l.last
}

func (l *List) First() *Item {
	return l.first
}

func (l *List) Len() uint {
	return l.len
}

func (l *List) PushFront(v interface{}) {
	l.first = NewItem(v, nil, l.first)
	if l.first.Prev() != nil {
		l.first.Prev().SetNext(l.first)
	}
	l.len++
	if l.len == 1 {
		l.last = l.first
	}
}

func (l *List) PushBack(v interface{}) {
	l.last = NewItem(v, l.last, nil)
	if l.last.Next() != nil {
		l.last.Next().SetPrev(l.first)
	}
	if l.len == 0 {
		l.first = l.last
	}
	l.len++
}

func (l *List) Remove(item *Item) error {
	if item == l.last {
		l.last = item.Next()
	} else if item == l.first {
		l.first = item.Prev()
	} else {
		if item.Next() == nil || item.Prev() == nil {
			return errors.New("list not contain this item")
		}
		item.Next().SetPrev(item.Prev())
		item.Prev().SetNext(item.Next())
	}
	l.len--
	return nil
}
