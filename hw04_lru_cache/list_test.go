package hw04lrucache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := NewList()

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("check PushFront: push item to empty list", func(t *testing.T) {
		l := NewList()

		l.PushFront("Red")
		require.Equal(t, 1, l.Len())
		require.Equal(t, "Red", l.Front().Value)
		require.Equal(t, "Red", l.Back().Value)
	})

	t.Run("check PushFront: push items to list", func(t *testing.T) {
		l := NewList()

		l.PushFront("Henry")
		l.PushFront(98)
		l.PushFront(true)
		require.Equal(t, 3, l.Len())
		require.Equal(t, true, l.Front().Value)
		require.Equal(t, "Henry", l.Back().Value)
	})

	t.Run("check PushBack: push item to empty list", func(t *testing.T) {
		l := NewList()

		l.PushBack(64)
		require.Equal(t, 1, l.Len())
		require.Equal(t, 64, l.Front().Value)
		require.Equal(t, 64, l.Back().Value)
	})

	t.Run("check PushBack: push items to lists", func(t *testing.T) {
		l := NewList()

		l.PushBack(4)
		l.PushBack(34)
		l.PushBack("4")
		l.PushBack(51)
		require.Equal(t, 4, l.Len())
		require.Equal(t, 4, l.Front().Value)
		require.Equal(t, 51, l.Back().Value)
	})

	t.Run("check Remove front item", func(t *testing.T) {
		l := NewList()

		l.PushFront(6)
		l.PushBack(18.9)
		l.PushBack(92)
		require.Equal(t, 3, l.Len())

		front := l.Front()
		l.Remove(front)

		require.Equal(t, 2, l.Len())
		require.Equal(t, 18.9, l.Front().Value)
		require.Equal(t, 92, l.Back().Value)
	})

	t.Run("check Remove middle item", func(t *testing.T) {
		l := NewList()

		l.PushFront(6)
		l.PushBack(18)
		l.PushBack(false)
		require.Equal(t, 3, l.Len())

		middle := l.Front().Next
		l.Remove(middle)

		require.Equal(t, 2, l.Len())
		require.Equal(t, 6, l.Front().Value)
		require.Equal(t, false, l.Back().Value)
	})

	t.Run("check Remove back item", func(t *testing.T) {
		l := NewList()

		l.PushFront(6.02)
		l.PushBack(18)
		l.PushBack(92.65)
		require.Equal(t, 3, l.Len())

		back := l.Back()
		l.Remove(back)

		require.Equal(t, 2, l.Len())
		require.Equal(t, 6.02, l.Front().Value)
		require.Equal(t, 18, l.Back().Value)
	})

	t.Run("check Remove item from one element list", func(t *testing.T) {
		l := NewList()

		l.PushFront(70209)
		require.Equal(t, 1, l.Len())

		item := l.Front()
		l.Remove(item)

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("check MoveToFront: one item list", func(t *testing.T) {
		l := NewList()
		l.PushFront(69)

		l.MoveToFront(l.Front())

		require.Equal(t, 1, l.Len())
		require.Equal(t, 69, l.Front().Value)
		require.Equal(t, 69, l.Back().Value)
	})

	t.Run("check MoveToFront: two items in list", func(t *testing.T) {
		l := NewList()
		l.PushFront(69)
		l.PushBack("one")

		l.MoveToFront(l.Back())

		require.Equal(t, 2, l.Len())
		require.Equal(t, "one", l.Front().Value)
		require.Equal(t, 69, l.Back().Value)
	})

	t.Run("check MoveToFront: many items in list", func(t *testing.T) {
		l := NewList()
		l.PushFront(69)
		l.PushBack(265)
		l.PushFront(34)
		l.PushBack(2) //34 69 265 2

		l.MoveToFront(l.Back().Prev) // 265

		require.Equal(t, 4, l.Len())
		require.Equal(t, 265, l.Front().Value)
		require.Equal(t, 2, l.Back().Value)
	})

	t.Run("multiple operations", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]
		l.PushBack(30)  // [10, 20, 30]
		require.Equal(t, 3, l.Len())

		middle := l.Front().Next // 20
		l.Remove(middle)         // [10, 30]
		require.Equal(t, 2, l.Len())

		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		} // [80, 60, 40, 10, 30, 50, 70]

		require.Equal(t, 7, l.Len())
		require.Equal(t, 80, l.Front().Value)
		require.Equal(t, 70, l.Back().Value)

		l.MoveToFront(l.Front()) // [80, 60, 40, 10, 30, 50, 70]
		l.MoveToFront(l.Back())  // [70, 80, 60, 40, 10, 30, 50]

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{70, 80, 60, 40, 10, 30, 50}, elems)
	})
}
