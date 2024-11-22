package set_test

import (
	"testing"

	set "github.com/paskozdilar/go-set"
)

func TestContains(t *testing.T) {
	s := make(map[int]string)
	e := 420

	if set.Contains(s, e) {
		t.Errorf("not inserted element found: %d", e)
	}
}

func TestInsert(t *testing.T) {
	s := make(map[int]string)
	e := 420

	set.Insert(s, e)
	if !set.Contains(s, e) {
		t.Errorf("inserted element not found: %d", e)
	}
}

func TestRemove(t *testing.T) {
	s := make(set.Set[int])
	e := 420

	set.Insert(s, e)
	set.Remove(s, e)
	if set.Contains(s, e) {
		t.Errorf("removed element found: %d", e)
	}
}

func TestUnion(t *testing.T) {
	s1 := make(set.Set[int])
	s2 := make(set.Set[int])
	s3 := make(set.Set[int])

	for i := 0; i < 10; i++ {
		s1.Insert(i)     // 0..9
		s2.Insert(i + 3) // 3..12
		s3.Insert(i + 6) // 6..15
	}

	su := s1.Union(s2, s3)

	if len(su) != 16 {
		t.Errorf("union of size %d; expected %d", len(su), 16)
	}

	for i := 0; i < 16; i++ {
		if !set.Contains(su, i) {
			t.Errorf("element not found: %d", i)
		}
	}
}

func TestIntersection(t *testing.T) {
	s1 := make(set.Set[int])
	s2 := make(set.Set[int])
	s3 := make(set.Set[int])

	for i := 0; i < 10; i++ {
		s1.Insert(i)     // 0..9
		s2.Insert(i + 3) // 3..12
		s3.Insert(i + 6) // 6..15
	}

	si := s1.Intersection(s2, s3)

	if len(si) != 4 {
		t.Errorf("union of size %d; expected %d", len(si), 4)
	}

	for i := 6; i < 10; i++ {
		if !set.Contains(si, i) {
			t.Errorf("element not found: %d", i)
		}
	}
}

func TestSymmetricDifference(t *testing.T) {
	s1 := make(set.Set[int])
	s2 := make(set.Set[int])
	s3 := make(set.Set[int])

	for i := 0; i < 10; i++ {
		s1.Insert(i)     // [0..9]
		s2.Insert(i + 3) // [3..12]
		s3.Insert(i + 6) // [6..15]
		//  0  1  2  3  4  5  6  7  8  9
		//           3  4  5  6  7  8  9 10 11 12
		//                    6  7  8  9 10 11 12 13 14 15
		// -----------------------------------------------
		//  0  1  2           6  7  8  9          13 14 15
	}

	ssd := s1.SymmetricDifference(s2, s3)

	if len(ssd) != 3+4+3 {
		t.Errorf("union of size %d; expected %d", len(ssd), 3+4+3)
	}

	for i := 0; i < 3; i++ {
		if !set.Contains(ssd, i) {
			t.Errorf("element not found: %d", i)
		}
	}

	for i := 6; i < 10; i++ {
		if !set.Contains(ssd, i) {
			t.Errorf("element not found: %d", i)
		}
	}

	for i := 13; i < 16; i++ {
		if !set.Contains(ssd, i) {
			t.Errorf("element not found: %d", i)
		}
	}
}
