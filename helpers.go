package collection

// equalSet checks if tw slices are equal sets or
// have the same unique unordered elements.
func equalSet[T comparable](a []T, b []T) bool {
	var (
		mapA = make(map[T]struct{}, len(a))
		mapB = make(map[T]struct{}, len(b))
		ok   bool
	)

	for _, v := range a {
		mapA[v] = struct{}{}
	}
	for _, v := range b {
		mapB[v] = struct{}{}
	}

	for k := range mapA {
		if _, ok = mapB[k]; !ok {
			return false
		}
	}

	for k := range mapB {
		if _, ok = mapA[k]; !ok {
			return false
		}
	}

	return true
}
