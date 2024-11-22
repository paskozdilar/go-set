package set

type Set[E comparable] map[E]struct{}

func New[E comparable]() Set[E] {
	return make(Set[E])
}

func (s Set[E]) Contains(e E) bool {
	return Contains(s, e)
}

func (s Set[E]) Insert(e E) {
	Insert(s, e)
}

func (s Set[E]) Remove(e E) {
	Remove(s, e)
}

func (s Set[T]) Union(sets ...Set[T]) Set[T] {
	return Union(append(sets, s)...)
}

func (s Set[T]) Intersection(sets ...Set[T]) Set[T] {
	return Intersection(append(sets, s)...)
}

func (s Set[T]) SymmetricDifference(sets ...Set[T]) Set[T] {
	return SymmetricDifference(append(sets, s)...)
}

func (s Set[T]) Elements() []T {
	return Elements(s)
}
