package collection_test

import (
	"sort"
	"testing"

	"slices"

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

func TestMapFirst(t *testing.T) {
	cases := []struct {
		name      string
		source    map[string]int
		predicate func(string, int) bool
		want      collection.KV[string, int]
		wantOk    bool
	}{
		{
			"Found matching element",
			map[string]int{"a": 1, "b": 2, "c": 3},
			func(k string, v int) bool { return v%2 == 0 },
			collection.KV[string, int]{Key: "b", Value: 2},
			true,
		},
		{
			"No matching element",
			map[string]int{"a": 1, "b": 3, "c": 5},
			func(k string, v int) bool { return v%2 == 0 },
			collection.KV[string, int]{},
			false,
		},
		{
			"Empty map",
			map[string]int{},
			func(k string, v int) bool { return v > 0 },
			collection.KV[string, int]{},
			false,
		},
		{
			"Key-based predicate",
			map[string]int{"apple": 1, "banana": 2, "cherry": 3},
			func(k string, v int) bool { return k == "banana" },
			collection.KV[string, int]{Key: "banana", Value: 2},
			true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, gotOk := collection.MapFirst(tc.source, tc.predicate)

			if gotOk != tc.wantOk {
				t.Errorf("MapFirst(%v, predicate) ok = %v; want %v", tc.source, gotOk, tc.wantOk)
				return
			}

			if got != tc.want {
				t.Errorf("MapFirst(%v, predicate) = (%v, %v); want (%v, %v)", tc.source, got, gotOk, tc.want, tc.wantOk)
			}
		})
	}
}

func TestMapFirstMultipleMatches(t *testing.T) {
	source := map[string]int{"a": 2, "b": 4, "c": 6}
	predicate := func(k string, v int) bool { return v%2 == 0 }
	
	got, ok := collection.MapFirst(source, predicate)
	
	if !ok {
		t.Errorf("MapFirst should return true for multiple matches")
		return
	}
	
	// Verify the returned value exists in the source map
	sourceValue, exists := source[got.Key]
	if !exists {
		t.Errorf("MapFirst returned key %q that doesn't exist in source map", got.Key)
		return
	}
	
	// Verify the returned value matches the source
	if got.Value != sourceValue {
		t.Errorf("MapFirst returned value %v for key %q; want %v", got.Value, got.Key, sourceValue)
		return
	}
	
	// Verify the returned pair satisfies the predicate
	if !predicate(got.Key, got.Value) {
		t.Errorf("MapFirst returned key-value pair (%q, %v) that doesn't match predicate", got.Key, got.Value)
	}
}

func TestMapFirstAnyElement(t *testing.T) {
	source := map[int]string{1: "one", 2: "two", 3: "three"}
	
	got, ok := collection.MapFirst(source, func(k int, v string) bool { return true })
	
	if !ok {
		t.Errorf("MapFirst should return true for any element predicate")
		return
	}
	
	// Verify the returned element exists in source
	sourceValue, exists := source[got.Key]
	if !exists {
		t.Errorf("MapFirst returned key %d that doesn't exist in source map", got.Key)
		return
	}
	
	if got.Value != sourceValue {
		t.Errorf("MapFirst returned value %q for key %d; want %q", got.Value, got.Key, sourceValue)
	}
}

func TestMapFirstNilMap(t *testing.T) {
	var source map[string]int
	
	got, ok := collection.MapFirst(source, func(k string, v int) bool { return true })
	
	if ok {
		t.Errorf("MapFirst should return false for nil map, got ok=%v", ok)
	}
	
	// Verify zero value returned
	if got.Key != "" || got.Value != 0 {
		t.Errorf("MapFirst should return zero values for nil map, got %+v", got)
	}
}

func TestMapFirstExistsOnly(t *testing.T) {
	source := map[string]int{"a": 5, "b": 15, "c": 25}
	
	// Test existence check - only care about ok value
	_, exists := collection.MapFirst(source, func(k string, v int) bool { return v > 10 })
	if !exists {
		t.Errorf("MapFirst should find elements > 10, got exists=%v", exists)
	}
	
	// Test non-existence
	_, notExists := collection.MapFirst(source, func(k string, v int) bool { return v > 100 })
	if notExists {
		t.Errorf("MapFirst should not find elements > 100, got exists=%v", notExists)
	}
}

type CustomKey struct {
	ID   int
	Name string
}

type CustomValue struct {
	Data  string
	Count int
}

func TestMapFirstCustomTypes(t *testing.T) {
	source := map[CustomKey]CustomValue{
		{ID: 1, Name: "first"}:  {Data: "data1", Count: 10},
		{ID: 2, Name: "second"}: {Data: "data2", Count: 20},
		{ID: 3, Name: "third"}:  {Data: "data3", Count: 30},
	}
	
	got, ok := collection.MapFirst(source, func(k CustomKey, v CustomValue) bool {
		return v.Count > 15
	})
	
	if !ok {
		t.Errorf("MapFirst should find element with Count > 15")
		return
	}
	
	// Verify the returned element exists and matches predicate
	sourceValue, exists := source[got.Key]
	if !exists {
		t.Errorf("MapFirst returned key %+v that doesn't exist in source map", got.Key)
		return
	}
	
	if got.Value != sourceValue {
		t.Errorf("MapFirst returned value %+v; want %+v", got.Value, sourceValue)
		return
	}
	
	if got.Value.Count <= 15 {
		t.Errorf("MapFirst returned element with Count=%d, should be > 15", got.Value.Count)
	}
}


