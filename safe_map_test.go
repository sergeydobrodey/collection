package collection

import (
	"reflect"
	"sort"
	"testing"
)

func TestSafeMap(t *testing.T) {
	tests := []struct {
		name   string
		action func(m *SafeMap[string, int]) interface{}
		expect interface{}
	}{
		{
			name: "Set and Get",
			action: func(m *SafeMap[string, int]) interface{} {
				m.Set("key1", 100)
				val, ok := m.Get("key1")
				return struct {
					value int
					ok    bool
				}{val, ok}
			},
			expect: struct {
				value int
				ok    bool
			}{100, true},
		},
		{
			name: "Get non-existent key",
			action: func(m *SafeMap[string, int]) interface{} {
				_, ok := m.Get("keyDoesNotExist")
				return ok
			},
			expect: false,
		},
		{
			name: "Delete key",
			action: func(m *SafeMap[string, int]) interface{} {
				m.Set("keyToDelete", 42)
				m.Delete("keyToDelete")
				_, ok := m.Get("keyToDelete")
				return ok
			},
			expect: false,
		},
		{
			name: "Has key",
			action: func(m *SafeMap[string, int]) interface{} {
				m.Set("existingKey", 1)
				return m.Has("existingKey")
			},
			expect: true,
		},
		{
			name: "Len",
			action: func(m *SafeMap[string, int]) interface{} {
				m.Set("k1", 1)
				m.Set("k2", 2)
				return m.Len()
			},
			expect: 2,
		},
		{
			name: "Clear map",
			action: func(m *SafeMap[string, int]) interface{} {
				m.Set("k1", 1)
				m.Set("k2", 2)
				m.Clear()
				return m.Len()
			},
			expect: 0,
		},
		{
			name: "Keys method",
			action: func(m *SafeMap[string, int]) interface{} {
				m.Set("k1", 1)
				m.Set("k2", 2)
				return m.Keys()
			},
			expect: []string{"k1", "k2"},
		},
		{
			name: "Values method",
			action: func(m *SafeMap[string, int]) interface{} {
				m.Set("k1", 1)
				m.Set("k2", 2)
				return m.Values()
			},
			expect: []int{1, 2},
		},
		{
			name: "ForEach method",
			action: func(m *SafeMap[string, int]) interface{} {
				results := make(map[string]int)
				m.Set("k1", 1)
				m.Set("k2", 2)
				m.ForEach(func(k string, v int) {
					results[k] = v
				})
				return results
			},
			expect: map[string]int{"k1": 1, "k2": 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewSafeMap[string, int]()
			result := tt.action(m)

			if !compareResults(result, tt.expect) {
				t.Errorf("expected %v, got %v", tt.expect, result)
			}
		})
	}
}

func compareResults(a, b interface{}) bool {
	switch a := a.(type) {
	case []string:
		b := b.([]string)
		sort.Strings(a)
		sort.Strings(b)
		return reflect.DeepEqual(a, b)
	case []int:
		b := b.([]int)
		sort.Ints(a)
		sort.Ints(b)
		return reflect.DeepEqual(a, b)
	case map[string]int:
		b := b.(map[string]int)
		return reflect.DeepEqual(a, b)
	case struct {
		value int
		ok    bool
	}:
		b := b.(struct {
			value int
			ok    bool
		})
		return a.value == b.value && a.ok == b.ok
	default:
		return a == b
	}
}
