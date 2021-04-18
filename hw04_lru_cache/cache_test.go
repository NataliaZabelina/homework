package hw04lrucache

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type cacheSuite struct {
	suite.Suite
}

func (s *cacheSuite) TestGetFromEmptyCache() {
	c := NewCache(10)

	_, ok := c.Get("aaa")
	s.False(ok)

	_, ok = c.Get("bbb")
	s.False(ok)
}

func (s *cacheSuite) TestGetAndSetCheckCache() {
	c := NewCache(5)

	wasInCache := c.Set("aaa", 100)
	s.False(wasInCache)

	wasInCache = c.Set("bbb", 200)
	s.False(wasInCache)

	val, ok := c.Get("aaa")
	s.True(ok)
	s.Equal(100, val)

	val, ok = c.Get("bbb")
	s.True(ok)
	s.Equal(200, val)

	wasInCache = c.Set("aaa", 300)
	s.True(wasInCache)

	val, ok = c.Get("aaa")
	s.True(ok)
	s.Equal(300, val)

	val, ok = c.Get("ccc")
	s.False(ok)
	s.Nil(val)
}

func (s *cacheSuite) TestCheckCapacityCache() {
	c := NewCache(3)
	c.Set("item_1", 78)
	c.Set("item_2", 73)
	c.Set("item_3", 12)

	val, ok := c.Get("item_1")
	s.Equal(78, val)
	s.True(ok)

	val, ok = c.Get("item_2")
	s.Equal(73, val)
	s.True(ok)

	val, ok = c.Get("item_3")
	s.Equal(12, val)
	s.True(ok)

	c.Set("item_4", 120)
	val, ok = c.Get("item_4")
	s.Equal(120, val)
	s.True(ok)

	val, ok = c.Get("item_1")
	s.False(ok)
	s.Nil(val)
}

func (s *cacheSuite) TestRemoveItemsDueToCapacityCache() {
	c := NewCache(2)
	c.Set("item_1", "ONE")
	c.Set("item_2", "TWO")

	val, ok := c.Get("item_1")
	s.Equal("ONE", val)
	s.True(ok)

	val, ok = c.Get("item_2")
	s.Equal("TWO", val)
	s.True(ok)

	c.Set("item_3", "THREE")
	val, ok = c.Get("item_3")
	s.Equal("THREE", val)
	s.True(ok)

	val, ok = c.Get("item_1")
	s.False(ok)
	s.Nil(val)

	c.Set("item_4", "FOUR")
	val, ok = c.Get("item_4")
	s.Equal("FOUR", val)
	s.True(ok)

	val, ok = c.Get("item_2")
	s.False(ok)
	s.Nil(val)
}

func (s *cacheSuite) TestClearCache() {
	c := NewCache(2)
	c.Set("item_1", true)
	c.Set("item_2", false)

	val, ok := c.Get("item_1")
	s.Equal(true, val)
	s.True(ok)

	val, ok = c.Get("item_2")
	s.Equal(false, val)
	s.True(ok)

	c.Clear()

	val, ok = c.Get("item_1")
	s.False(ok)
	s.Nil(val)

	val, ok = c.Get("item_2")
	s.False(ok)
	s.Nil(val)
}

func (s *cacheSuite) TestCheckLruLogic() {
	c := NewCache(3)
	c.Set("item_1", 6)
	c.Set("item_2", 8)
	c.Set("item_3", 2)

	val, ok := c.Get("item_1")
	s.Equal(6, val)
	s.True(ok)

	val, ok = c.Get("item_2")
	s.Equal(8, val)
	s.True(ok)

	val, ok = c.Get("item_3")
	s.Equal(2, val)
	s.True(ok)

	c.Set("item_2", 90)
	_, ok = c.Get("item_3")
	s.True(ok)
	c.Set("item_2", 88)
	c.Set("item_3", 22)

	c.Set("item_4", "FOUR")
	val, ok = c.Get("item_4")
	s.Equal("FOUR", val)
	s.True(ok)

	val, ok = c.Get("item_1")
	s.False(ok)
	s.Nil(val)

	val, ok = c.Get("item_2")
	s.Equal(88, val)
	s.True(ok)

	val, ok = c.Get("item_3")
	s.Equal(22, val)
	s.True(ok)
}

func TestCacheSuite(t *testing.T) {
	suite.Run(t, new(cacheSuite))
}
