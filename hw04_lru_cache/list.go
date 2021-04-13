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
	Length    int
	FrontElem *ListItem
	BackElem  *ListItem
}

func (l *list) Len() int {
	return l.Length
}

func (l *list) Front() *ListItem {
	return l.FrontElem
}

func (l *list) Back() *ListItem {
	return l.BackElem
}

func (l *list) pushAfter(item *ListItem, newItem *ListItem) {
	newItem.Prev = item
	if item.Next == nil {
		newItem.Next = nil
		l.BackElem = newItem
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
		l.FrontElem = newItem
	} else {
		newItem.Prev = item.Prev
		item.Prev.Next = newItem
	}

	item.Prev = newItem
}

func (l *list) PushFront(v interface{}) *ListItem {
	newItem := &ListItem{Value: v}

	if l.FrontElem == nil {
		l.FrontElem = newItem
		l.BackElem = newItem
		newItem.Prev = nil
		newItem.Next = nil
	} else {
		l.pushBefore(l.FrontElem, newItem)
	}

	l.Length++

	return newItem
}

func (l *list) PushBack(v interface{}) *ListItem {
	newItem := &ListItem{Value: v}

	if l.BackElem == nil {
		l.PushFront(v)
	} else {
		l.pushAfter(l.BackElem, newItem)
		l.Length++
	}

	return newItem
}

func (l *list) Remove(item *ListItem) {
	if item.Prev == nil {
		l.FrontElem = item.Next
	} else {
		item.Prev.Next = item.Next
	}
	if item.Next == nil {
		l.BackElem = item.Prev
	} else {
		item.Next.Prev = item.Prev
	}

	l.Length--
}

func (l *list) MoveToFront(item *ListItem) {
	l.Remove(item)
	l.PushFront(item.Value)
}

func NewList() List {
	return new(list)
}
