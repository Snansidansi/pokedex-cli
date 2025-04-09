package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const intervall = 10 * time.Second
	cases := []struct {
		key   string
		value []byte
	}{
		{
			key:   "www.test.com",
			value: []byte("Testdata"),
		},
		{
			key:   "www.test.com/testing",
			value: []byte("Testdata - testing"),
		},
		{
			key:   "www.example.com",
			value: []byte("Some examples"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Subtest %v", i), func(t *testing.T) {
			cache := NewCache(intervall, intervall)
			cache.Add(c.key, c.value)

			val, ok := cache.Get(c.key)
			if !ok {
				t.Error("expected cache to contain key")
				return
			}
			if string(val) != string(c.value) {
				t.Error("cache does not contain expected value for key")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const (
		maxCacheAge   = 4 * time.Millisecond
		checkInterval = 2 * time.Millisecond
		sleepTime     = 10 * time.Millisecond
		key           = "www.test.com"
	)

	cache := NewCache(checkInterval, maxCacheAge)
	cache.Add(key, []byte("test"))

	_, ok := cache.Get(key)
	if !ok {
		t.Error("cache does not contain added key")
		return
	}

	time.Sleep(sleepTime)

	_, ok = cache.Get(key)
	if ok {
		t.Errorf("key should be removed from cache")
		return
	}
}
