package set

// Union returns a new set containing all the values in both a and b
func Union(a, b Set) Set {
	s := NewHashSet()
	s.AddAll(a.Slice()...)
	s.AddAll(b.Slice()...)
	return s
}

// Intersection returns a new set containing the values that exist in both a
// and b
func Intersection(a, b Set) Set {
	s := NewHashSet()
	aSlice := a.Slice()
	for _, item := range aSlice {
		if b.Contains(item) {
			s.Add(item)
		}
	}

	bSlice := a.Slice()
	for _, item := range bSlice {
		if a.Contains(item) {
			s.Add(item)
		}
	}

	return s
}

// Complement returns a new set containing all values that exist in a after
// subtracting the values in b
func Complement(a, b Set) Set {
	s := NewHashSet()
	aSlice := s.Slice()
	for _, item := range aSlice {
		if !b.Contains(item) {
			s.Add(item)
		}
	}

	return s
}
