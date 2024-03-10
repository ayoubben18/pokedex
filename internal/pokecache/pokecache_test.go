package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	cache := NewCache(time.Millisecond)
	if cache.cache == nil {
		t.Error("cache is nil")
	}
}

func TestAddToCache(t *testing.T) {
	cache := NewCache(time.Millisecond)
	cache.Add("key1", []byte("value1"))

	if len(cache.cache) != 1 {
		t.Errorf("expected 1 item in cache, got %d", len(cache.cache))
	}
}

func TestGetFromCache(t *testing.T) {
	cache := NewCache(time.Millisecond)
	cases := []struct {
		key      string
		expected []byte
	}{
		{key: "key1", expected: []byte("value1")},
		{key: "key2", expected: nil},
		//test them
	}
	for _, cs := range cases {
		cache.Add(cs.key, cs.expected)
		val, ok := cache.Get(cs.key)
		if !ok {
			t.Errorf("key %s not found in cache", cs.key)
		}
		if string(val) != string(cs.expected) {
			t.Errorf("expected value %s, got %s", string(cs.expected), string(val))
		}
	}
}

func TestReapCache(t *testing.T) {
	cache := NewCache(10 * time.Millisecond)
	cache.Add("key1", []byte("value1"))
	time.Sleep(20 * time.Millisecond)
	_, ok := cache.Get("key1")
	if ok {
		t.Errorf("key1 should have been deleted from cache")
	}
	cache.Add("key2", []byte("value2"))
	_, ok = cache.Get("key2")
	if !ok {
		t.Errorf("key2 should have been added to cache")
	}
}

func TestReapCacheFail(t *testing.T) {
	cache := NewCache(10 * time.Millisecond)
	cache.Add("key1", []byte("value1"))
	_, ok := cache.Get("key1")
	if !ok {
		t.Errorf("key1 should have been added to cache")
	}
	time.Sleep(20 * time.Millisecond)
	_, ok = cache.Get("key1")
	if ok {
		t.Errorf("key1 should have been deleted from cache")
	}
}
