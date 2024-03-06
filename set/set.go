package set

// Set implements the set data structure, for holding unique values.
// Adding elements to a nil Set will panic.
type Set[K comparable] map[K]struct{}

// New returns a new Set, populated with optionally provided values.
func New[K comparable](ks ...K) Set[K] {
	new := Set[K]{}
	for _, k := range ks {
		new[k] = struct{}{}
	}
	return new
}

func (s Set[K]) ToSlice() []K {
	if s == nil {
		return nil
	}
	out := make([]K, 0, len(s))
	for k := range s {
		out = append(out, k)
	}
	return out
}

func (s Set[K]) Add(k K) {
	s[k] = struct{}{}
}

func (s Set[K]) AddAll(ks ...K) {
	for _, k := range ks {
		s.Add(k)
	}
}

func (s Set[K]) Remove(k K) {
	if s == nil {
		return
	}
	delete(s, k)
}

func (s Set[K]) Size() int {
	if s == nil {
		return 0
	}
	return len(s)
}

func (s Set[K]) Contains(k K) bool {
	if s == nil {
		return false
	}
	_, ok := s[k]
	return ok
}

func (s Set[K]) IsEmpty() bool {
	return len(s) == 0
}
