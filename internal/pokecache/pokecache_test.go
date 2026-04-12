package pokecache

import (
	"testing"
	"time"
)

// test NewCahce() function
func TestCreateCache(t *testing.T) {
	cache_obj := NewCache(time.Millisecond)
	if cache_obj.cache == nil {
		t.Errorf("Cache is nil")
	}
}

// test Add() and Get() method
func TestAddGetCache(t *testing.T) {
	cache_obj := NewCache(time.Millisecond)

	test_cases := []struct {
		inputKey string
		inputVal []byte
	}{
		{
			inputKey: "key1",
			inputVal: []byte("val1"),
		},
		{
			inputKey: "key2",
			inputVal: []byte("val2"),
		},
		{
			inputKey: "",
			inputVal: []byte("val3"),
		},
	}

	for _, test_case := range test_cases {
		cache_obj.Add(test_case.inputKey, test_case.inputVal)

		actual, ok := cache_obj.Get(test_case.inputKey)

		if !ok {
			t.Errorf("%s not found", test_case.inputKey)
			continue
		}
		if string(actual) != string(test_case.inputVal) {
			t.Errorf("value doesn't match, actual: %v, expected: %s", string(actual), string(test_case.inputVal))
		}
	}
}

// test reap() and reapLoop() method with interval time
func TestReapCache(t *testing.T) {
	interval := time.Millisecond * 10

	keyOne := "key1"

	cache_obj := NewCache(interval)
	cache_obj.Add(keyOne, []byte("val1"))

	// we expect to find the key here since we haven't passed interval time yet
	_, ok := cache_obj.Get(keyOne)
	if !ok {
		t.Errorf("%s should not have been reaped yet", keyOne)
	}

	time.Sleep(interval + time.Millisecond)

	// we expect to not find key here since we have already passed interval time
	_, ok = cache_obj.Get(keyOne)
	if ok {
		t.Errorf("%s should have been reaped", keyOne)
	}
}
