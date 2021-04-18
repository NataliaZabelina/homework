package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	length    int
	frontElem *ListItem
	backElem  *ListItem
}

func (l *list) Len() int {
	return l.length
}

func (l *list) Front() *ListItem {
	return l.frontElem
}

func (l *list) Back() *ListItem {
	return l.backElem
}

func (l *list) pushAfter(item *ListItem, newItem *ListItem) {
	newItem.Prev = item
	if item.Next == nil {
		newItem.Next = nil
		l.backElem = newItem
	} else {
		newItem.Next = item.Next
		item.Next.Prev = newItem
	}
	item.Next = newItem
}

func (l *list) pushBefore(item *ListItem, newItem *ListItem) {
	newItem.Next = item
	if item.Prev == nil {
		newItem.Prev = nil
		l.frontElem = newItem
	} else {
		newItem.Prev = item.Prev
		item.Prev.Next = newItem
	}

	item.Prev = newItem
}

func (l *list) PushFront(v interface{}) *ListItem {
	newItem := &ListItem{Value: v}

	if l.frontElem == nil {
		l.frontElem = newItem
		l.backElem = newItem
		newItem.Prev = nil
		newItem.Next = nil
	} else {
		l.pushBefore(l.frontElem, newItem)
	}

	l.length++

	return newItem
}

func (l *list) PushBack(v interface{}) *ListItem {
	newItem := &ListItem{Value: v}

	if l.backElem == nil {
		l.PushFront(v)
	} else {
		l.pushAfter(l.backElem, newItem)
		l.length++
	}

	return newItem
}

func (l *list) Remove(item *ListItem) {
	if item.Prev == nil {
		l.frontElem = item.Next
	} else {
		item.Prev.Next = item.Next
	}
	if item.Next == nil {
		l.backElem = item.Prev
	} else {
		item.Next.Prev = item.Prev
	}

	l.length--
}

func (l *list) MoveToFront(item *ListItem) {
	l.Remove(item)
	l.PushFront(item.Value)
}

func NewList() List {
	return new(list)
}
