package hw04lrucache

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type CacheSuite struct {
	suite.Suite
}

func (s *CacheSuite) TestGetFromEmptyCache() {
	t := s.T()
	c := NewCache(10)

	_, ok := c.Get("aaa")
	require.False(t, ok)

	_, ok = c.Get("bbb")
	require.False(t, ok)
}

func (s *CacheSuite) TestGetAndSetCheckCache() {
	t := s.T()
	c := NewCache(5)

	wasInCache := c.Set("aaa", 100)
	require.False(t, wasInCache)

	wasInCache = c.Set("bbb", 200)
	require.False(t, wasInCache)

	val, ok := c.Get("aaa")
	require.True(t, ok)
	require.Equal(t, 100, val)

	val, ok = c.Get("bbb")
	require.True(t, ok)
	require.Equal(t, 200, val)

	wasInCache = c.Set("aaa", 300)
	require.True(t, wasInCache)

	val, ok = c.Get("aaa")
	require.True(t, ok)
	require.Equal(t, 300, val)

	val, ok = c.Get("ccc")
	require.False(t, ok)
	require.Nil(t, val)
}

func (s *CacheSuite) TestCheckCapacityCache() {
	t := s.T()
	c := NewCache(3)
	c.Set("item_1", 78)
	c.Set("item_2", 73)
	c.Set("item_3", 12)

	val, ok := c.Get("item_1")
	require.Equal(t, 78, val)
	require.True(t, ok)

	val, ok = c.Get("item_2")
	require.Equal(t, 73, val)
	require.True(t, ok)

	val, ok = c.Get("item_3")
	require.Equal(t, 12, val)
	require.True(t, ok)

	c.Set("item_4", 120)
	val, ok = c.Get("item_4")
	require.Equal(t, 120, val)
	require.True(t, ok)

	val, ok = c.Get("item_1")
	require.False(t, ok)
	require.Nil(t, val)
}

func (s *CacheSuite) TestRemoveItemsDueToCapacityCache() {
	t := s.T()
	c := NewCache(2)
	c.Set("item_1", "ONE")
	c.Set("item_2", "TWO")

	val, ok := c.Get("item_1")
	require.Equal(t, "ONE", val)
	require.True(t, ok)

	val, ok = c.Get("item_2")
	require.Equal(t, "TWO", val)
	require.True(t, ok)

	c.Set("item_3", "THREE")
	val, ok = c.Get("item_3")
	require.Equal(t, "THREE", val)
	require.True(t, ok)

	val, ok = c.Get("item_1")
	require.False(t, ok)
	require.Nil(t, val)

	c.Set("item_4", "FOUR")
	val, ok = c.Get("item_4")
	require.Equal(t, "FOUR", val)
	require.True(t, ok)

	val, ok = c.Get("item_2")
	require.False(t, ok)
	require.Nil(t, val)
}

func (s *CacheSuite) TestClearCache() {
	t := s.T()
	c := NewCache(2)
	c.Set("item_1", true)
	c.Set("item_2", false)

	val, ok := c.Get("item_1")
	require.Equal(t, true, val)
	require.True(t, ok)

	val, ok = c.Get("item_2")
	require.Equal(t, false, val)
	require.True(t, ok)

	c.Clear()

	val, ok = c.Get("item_1")
	require.False(t, ok)
	require.Nil(t, val)

	val, ok = c.Get("item_2")
	require.False(t, ok)
	require.Nil(t, val)
}

func (s *CacheSuite) TestCheckLruLogic() {
	t := s.T()
	c := NewCache(3)
	c.Set("item_1", 6)
	c.Set("item_2", 8)
	c.Set("item_3", 2)

	val, ok := c.Get("item_1")
	require.Equal(t, 6, val)
	require.True(t, ok)

	val, ok = c.Get("item_2")
	require.Equal(t, 8, val)
	require.True(t, ok)

	val, ok = c.Get("item_3")
	require.Equal(t, 2, val)
	require.True(t, ok)

	c.Set("item_2", 90)
	_, ok = c.Get("item_3")
	require.True(t, ok)
	c.Set("item_2", 88)
	c.Set("item_3", 22)

	c.Set("item_4", "FOUR")
	val, ok = c.Get("item_4")
	require.Equal(t, "FOUR", val)
	require.True(t, ok)

	val, ok = c.Get("item_1")
	require.False(t, ok)
	require.Nil(t, val)

	val, ok = c.Get("item_2")
	require.Equal(t, 88, val)
	require.True(t, ok)

	val, ok = c.Get("item_3")
	require.Equal(t, 22, val)
	require.True(t, ok)
}

func TestCacheSuite(t *testing.T) {
	suite.Run(t, new(CacheSuite))
}
