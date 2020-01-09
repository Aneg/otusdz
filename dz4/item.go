package dz4

func NewItem(value interface{}, Next *Item, Prev *Item) *Item {
	return &Item{
		value: value,
		next:  Next,
		prev:  Prev,
	}
}

type Item struct {
	value interface{}
	next  *Item
	prev  *Item
}

func (i *Item) SetPrev(item *Item) {
	i.prev = item
}

func (i *Item) SetNext(item *Item) {
	i.next = item
}

func (i *Item) Next() *Item {
	return i.next
}

func (i *Item) Prev() *Item {
	return i.prev
}

func (i *Item) Value() interface{} {
	return i.value
}
