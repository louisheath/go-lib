package set

// Set implements the set data structure, aliasing a map to an empty struct.
// It must be instantiated before use. Operting on a nil Set will panic, as would operating on a nil
// map.
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
	out := make([]K, 0, len(s))
	for k := range s {
		out = append(out, k)
	}
	return out
}

func (s Set[K]) Add(k K) {
	s[k] = struct{}{}
}

func (s Set[K]) AddAll(k []K) {
	for _, k := range k {
		s.Add(k)
	}
}

func (s Set[K]) Remove(k K) {
	delete(s, k)
}

func (s Set[K]) Size() int {
	return len(s)
}

func (s Set[K]) IsEmpty() bool {
	return len(s) == 0
}

func (s Set[K]) Contains(k K) bool {
	_, ok := s[k]
	return ok
}
