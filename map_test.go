package collection_test

import (
	"sort"
	"testing"

	"golang.org/x/exp/slices"

	"github.com/sergeydobrodey/collection"
)

func TestMapKeys(t *testing.T) {
	source := map[string]int{"a": 1, "b": 2, "c": 3}
	keys := collection.MapKeys(source)

	sort.Strings(keys)
	want := []string{"a", "b", "c"}

	if !slices.Equal(keys, want) {
		t.Errorf("MapKeys(%v) = %v; want %v", source, keys, want)
	}
}

func TestMapValues(t *testing.T) {
	source := map[string]int{"a": 1, "b": 2, "c": 3}
	values := collection.MapValues(source)

	sort.Ints(values)
	want := []int{1, 2, 3}

	if !slices.Equal(values, want) {
		t.Errorf("MapValues(%v) = %v; want %v", source, values, want)
	}
}

func initSyncMap(t *testing.T) *collection.SyncMap[int, string] {
	t.Helper()

	syncMap := collection.SyncMap[int, string]{}
	syncMap.Store(1, "one")
	syncMap.Store(2, "two")
	syncMap.Store(3, "three")
	return &syncMap
}

func TestSyncMapLoadOrStore(t *testing.T) {
	syncMap := initSyncMap(t)

	value, ok := syncMap.Load(2)
	if !ok || value != "two" {
		t.Errorf("Load(%v) = want %v, got %v", 2, "two", value)
	}

	actual, loaded := syncMap.LoadOrStore(4, "four")
	if loaded || actual != "four" {
		t.Errorf("LoadOrStore(%v) = want (%v, %v), got (%v, %v)", 4, "four", false, actual, loaded)
	}

	actual, loaded = syncMap.LoadOrStore(2, "new_two")
	if !loaded || actual != "two" {
		t.Errorf("LoadOrStore(%v) = want (%v, %v), got (%v, %v)", 2, "two", true, actual, loaded)
	}
}

func TestSyncMapLoadAndDelete(t *testing.T) {
	syncMap := initSyncMap(t)

	syncMap.Delete(3)
	if _, ok := syncMap.Load(3); ok {
		t.Errorf("Delete(%v) failed to delete key - key still exists", 3)
	}

	value, loaded := syncMap.LoadAndDelete(1)
	if !loaded || value != "one" {
		t.Errorf("LoadAndDelete(%v) = want (%v, %v), got (%v, %v)", 1, "one", true, value, loaded)
	}

	value, loaded = syncMap.LoadAndDelete(1)
	if loaded || value != "" {
		t.Errorf("LoadAndDelete(%v) = want (%v, %v), got (%v, %v)", 1, "", false, value, loaded)
	}

	if _, ok := syncMap.Load(1); ok {
		t.Errorf("LoadAndDelete(%v) failed to delete key - key still exists", 1)
	}
}

func TestSyncMapRange(t *testing.T) {
	syncMap := initSyncMap(t)

	var keys []int
	var values []string

	syncMap.Range(func(key int, value string) bool {
		keys = append(keys, key)
		values = append(values, value)
		return true
	})

	sort.Ints(keys)
	sort.Strings(values)

	wantKeys := []int{1, 2, 3}
	wantValues := []string{"one", "three", "two"}

	if !slices.Equal(keys, wantKeys) || !slices.Equal(values, wantValues) {
		t.Errorf("Range() = want (%v, %v), got (%v, %v)", wantKeys, wantValues, keys, values)
	}
}

func TestSyncMapCompareAndSwap(t *testing.T) {
	syncMap := initSyncMap(t)

	previous, loaded := syncMap.Swap(3, "new_three")
	if !loaded || previous != "three" {
		t.Errorf("Swap(%v) = want (%v, %v), got (%v, %v)", 3, "three", true, previous, loaded)
	}

	previous, loaded = syncMap.Swap(5, "new_five")
	if loaded || previous != "" {
		t.Errorf("Swap(%v) = want (%v, %v), got (%v, %v)", 5, "", false, previous, loaded)
	}

	if _, ok := syncMap.Load(5); !ok {
		t.Errorf("Swap(%v) failed to store key - value pair", 5)
	}

	if !syncMap.CompareAndSwap(2, "two", "updated_two") {
		t.Errorf("CompareAndSwap(%v) failed to swap value", 2)
	}

	if value, ok := syncMap.Load(2); !ok || value != "updated_two" {
		t.Errorf("CompareAndSwap(%v) failed to update value", 2)
	}

	if !syncMap.CompareAndDelete(2, "updated_two") {
		t.Errorf("CompareAndDelete(%v) failed to delete value", 2)
	}
}
