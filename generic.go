package set

type SetType[K comparable, V any] interface {
	~map[K]V
}

func Contains[S SetType[K, V], K comparable, V any](s S, k K) bool {
	_, ok := s[k]
	return ok
}

func Insert[S SetType[K, V], K comparable, V any](s S, k K) {
	var v V
	s[k] = v
}

func Remove[S SetType[K, V], K comparable, V any](s S, k K) {
	delete(s, k)
}

func Union[S SetType[K, V], K comparable, V any](sets ...S) S {
	union := make(S)

	for _, s := range sets {
		for k, v := range s {
			union[k] = v
		}
	}

	return union
}

func Intersection[S SetType[K, V], K comparable, V any](sets ...S) S {
	intersection := make(S)

	if len(sets) == 0 {
		return intersection
	}

	minIdx, minLen := 0, len(sets[0])
	for idx, s := range sets {
		if len(s) < minLen {
			minIdx, minLen = idx, len(s)
		}
	}
	minSet := sets[minIdx]

ITEM_LOOP:
	for k, v := range minSet {
		for _, s := range sets {
			if _, ok := s[k]; !ok {
				continue ITEM_LOOP
			}
		}
		intersection[k] = v
	}

	return intersection
}

func SymmetricDifference[S SetType[K, V], K comparable, V any](sets ...S) S {
	symdif := make(S)

	for k, v := range Union(sets...) {
		cnt := 0
		for _, s := range sets {
			if _, ok := s[k]; ok {
				cnt += 1
			}
		}
		if cnt%2 == 1 {
			symdif[k] = v
		}
	}

	return symdif
}

func Elements[S SetType[K, V], K comparable, V any](set S) []K {
	elements := make([]K, 0, len(set))

	for k := range set {
		elements = append(elements, k)
	}

	return elements
}
