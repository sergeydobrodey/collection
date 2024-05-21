package collection_test

import (
	"reflect"
	"sort"
	"testing"

	"github.com/sergeydobrodey/collection"
)

func TestMapKeys(t *testing.T) {
	source := map[string]int{"a": 1, "b": 2, "c": 3}
	keys := collection.MapKeys(source)

	sort.Strings(keys)
	expected := []string{"a", "b", "c"}
	if !reflect.DeepEqual(keys, expected) {
		t.Errorf("MapKeys returned %v, expected %v", keys, expected)
	}
}

func TestMapValues(t *testing.T) {
	source := map[string]int{"a": 1, "b": 2, "c": 3}
	values := collection.MapValues(source)

	sort.Ints(values)
	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(values, expected) {
		t.Errorf("MapValues returned %v, expected %v", values, expected)
	}
}

func TestSyncMap(t *testing.T) {
	syncMap := collection.SyncMap[int, string]{}
	syncMap.Store(1, "one")
	syncMap.Store(2, "two")
	syncMap.Store(3, "three")

	value, ok := syncMap.Load(2)
	if !ok || value != "two" {
		t.Errorf("Load returned incorrect value. Got %v, expected %v.", value, "two")
	}

	actual, loaded := syncMap.LoadOrStore(4, "four")
	if loaded || actual != "four" {
		t.Errorf("LoadOrStore returned incorrect value. Got (%v, %v), expected (%v, %v).", actual, loaded, "", false)
	}

	actual, loaded = syncMap.LoadOrStore(2, "new_two")
	if !loaded || actual != "two" {
		t.Errorf("LoadOrStore returned incorrect value. Got (%v, %v), expected (%v, %v).", actual, loaded, "two", true)
	}

	syncMap.Delete(3)
	if _, ok := syncMap.Load(3); ok {
		t.Errorf("Delete method failed to delete key. Key 3 still exists.")
	}

	value, loaded = syncMap.LoadAndDelete(1)
	if !loaded || value != "one" {
		t.Errorf("LoadAndDelete returned incorrect value. Got (%v, %v), expected (%v, %v).", value, loaded, "one", true)
	}

	value, loaded = syncMap.LoadAndDelete(1)
	if loaded || value != "" {
		t.Errorf("LoadAndDelete returned incorrect value. Got (%v, %v), expected (%v, %v).", value, loaded, "", false)
	}

	if _, ok := syncMap.Load(1); ok {
		t.Errorf("LoadAndDelete method failed to delete key. Key 1 still exists.")
	}

	keys := []int{}
	values := []string{}
	syncMap.Range(func(key int, value string) bool {
		keys = append(keys, key)
		values = append(values, value)
		return true
	})

	expectedKeys := []int{2, 4}
	expectedValues := []string{"four", "two"}
	sort.Ints(keys)
	sort.Strings(values)

	if !reflect.DeepEqual(expectedKeys, expectedKeys) || !reflect.DeepEqual(values, expectedValues) {
		t.Errorf("Range method returned incorrect result. Got (%v, %v), expected (%v, %v).", keys, values, expectedKeys, expectedValues)
	}

	previous, loaded := syncMap.Swap(5, "new_five")
	if loaded || previous != "" {
		t.Errorf("Swap method returned incorrect value. Got (%v, %v), expected (%v, %v).", previous, loaded, "", false)
	}

	previous, loaded = syncMap.Swap(4, "new_four")
	if !loaded || previous != "four" {
		t.Errorf("Swap method returned incorrect value. Got (%v, %v), expected (%v, %v).", previous, loaded, "four", true)
	}

	if _, ok := syncMap.Load(4); !ok {
		t.Errorf("Swap method failed to store new key-value pair.")
	}

	if !syncMap.CompareAndSwap(2, "two", "updated_two") {
		t.Errorf("CompareAndSwap method failed to update value.")
	}

	if value, ok := syncMap.Load(2); !ok || value != "updated_two" {
		t.Errorf("CompareAndSwap method failed to update value.")
	}

	// Test CompareAndDelete method.
	if !syncMap.CompareAndDelete(2, "updated_two") {
		t.Errorf("CompareAndDelete method failed to delete key-value pair.")
	}
}
