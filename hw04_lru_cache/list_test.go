package hw04lrucache

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type listSuite struct {
	suite.Suite
}

func (s *listSuite) TestEmptyList() {
	l := NewList()

	s.Equal(0, l.Len())
	s.Nil(l.Front())
	s.Nil(l.Back())
}

func (s *listSuite) TestPushFrontItemToEmptyList() {
	l := NewList()

	l.PushFront("Red")
	s.Equal(1, l.Len())
	s.Equal("Red", l.Front().Value)
	s.Equal("Red", l.Back().Value)
}

func (s *listSuite) TestPushFrontItemsToList() {
	l := NewList()

	l.PushFront("Henry")
	l.PushFront(98)
	l.PushFront(true)
	s.Equal(3, l.Len())
	s.Equal(true, l.Front().Value)
	s.Equal("Henry", l.Back().Value)
}

func (s *listSuite) TestPushBackItemToEmptyList() {
	l := NewList()

	l.PushBack(64)
	s.Equal(1, l.Len())
	s.Equal(64, l.Front().Value)
	s.Equal(64, l.Back().Value)
}

func (s *listSuite) TestRemoveFrontItemFromList() {
	l := NewList()

	l.PushFront(6)
	l.PushBack(18.9)
	l.PushBack(92)
	s.Equal(3, l.Len())

	front := l.Front()
	l.Remove(front)

	s.Equal(2, l.Len())
	s.Equal(18.9, l.Front().Value)
	s.Equal(92, l.Back().Value)
}

func (s *listSuite) TestPushBackItemsToList() {
	l := NewList()

	l.PushBack(4)
	l.PushBack(34)
	l.PushBack("4")
	l.PushBack(51)
	s.Equal(4, l.Len())
	s.Equal(4, l.Front().Value)
	s.Equal(51, l.Back().Value)
}

func (s *listSuite) TestRemoveMiddleItemFromList() {
	l := NewList()

	l.PushFront(6)
	l.PushBack(18)
	l.PushBack(false)
	s.Equal(3, l.Len())

	middle := l.Front().Next
	l.Remove(middle)

	s.Equal(2, l.Len())
	s.Equal(6, l.Front().Value)
	s.Equal(false, l.Back().Value)
}

func (s *listSuite) TestRemoveBackItemFromList() {
	l := NewList()

	l.PushFront(6.02)
	l.PushBack(18)
	l.PushBack(92.65)
	s.Equal(3, l.Len())

	back := l.Back()
	l.Remove(back)

	s.Equal(2, l.Len())
	s.Equal(6.02, l.Front().Value)
	s.Equal(18, l.Back().Value)
}

func (s *listSuite) TestRemoveFromOneItemList() {

	l := NewList()

	l.PushFront(70209)
	s.Equal(1, l.Len())

	item := l.Front()
	l.Remove(item)

	s.Equal(0, l.Len())
	s.Nil(l.Front())
	s.Nil(l.Back())
}

func (s *listSuite) TestMoveToFrontOneItemList() {
	l := NewList()
	l.PushFront(69)

	l.MoveToFront(l.Front())

	s.Equal(1, l.Len())
	s.Equal(69, l.Front().Value)
	s.Equal(69, l.Back().Value)
}

func (s *listSuite) TestMoveToFrontTwoItemsList() {
	l := NewList()
	l.PushFront(69)
	l.PushBack("one")

	l.MoveToFront(l.Back())

	s.Equal(2, l.Len())
	s.Equal("one", l.Front().Value)
	s.Equal(69, l.Back().Value)
}

func (s *listSuite) TestMoveToFrontManyItemsList() {
	l := NewList()
	l.PushFront(69)
	l.PushBack(265)
	l.PushFront(34)
	l.PushBack(2) // 34 69 265 2

	l.MoveToFront(l.Back().Prev) // 265

	s.Equal(4, l.Len())
	s.Equal(265, l.Front().Value)
	s.Equal(2, l.Back().Value)
}

func (s *listSuite) TestMultipleOperationsWithList() {
	l := NewList()

	l.PushFront(10) // [10]
	l.PushBack(20)  // [10, 20]
	l.PushBack(30)  // [10, 20, 30]
	s.Equal(3, l.Len())

	middle := l.Front().Next // 20
	l.Remove(middle)         // [10, 30]
	s.Equal(2, l.Len())

	for i, v := range [...]int{40, 50, 60, 70, 80} {
		if i%2 == 0 {
			l.PushFront(v)
		} else {
			l.PushBack(v)
		}
	} // [80, 60, 40, 10, 30, 50, 70]

	s.Equal(7, l.Len())
	s.Equal(80, l.Front().Value)
	s.Equal(70, l.Back().Value)

	l.MoveToFront(l.Front()) // [80, 60, 40, 10, 30, 50, 70]
	l.MoveToFront(l.Back())  // [70, 80, 60, 40, 10, 30, 50]

	elems := make([]int, 0, l.Len())
	for i := l.Front(); i != nil; i = i.Next {
		elems = append(elems, i.Value.(int))
	}
	s.Equal([]int{70, 80, 60, 40, 10, 30, 50}, elems)
}

func TestListSuite(t *testing.T) {
	suite.Run(t, new(listSuite))
}
