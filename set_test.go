package set_test

import (
	"testing"

	set "github.com/paskozdilar/go-set"
)

func TestSetContains(t *testing.T) {
	s := make(set.Set[int])
	e := 420

	if s.Contains(e) {
		t.Errorf("not inserted element found: %d", e)
	}
}

func TestSetInsert(t *testing.T) {
	s := make(set.Set[int])
	e := 420

	s.Insert(e)
	if !s.Contains(e) {
		t.Errorf("inserted element not found: %d", e)
	}
}

func TestSetRemove(t *testing.T) {
	s := make(set.Set[int])
	e := 420

	s.Insert(e)
	s.Remove(e)
	if s.Contains(e) {
		t.Errorf("removed element found: %d", e)
	}
}

func TestSetUnion(t *testing.T) {
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
		if !su.Contains(i) {
			t.Errorf("element not found: %d", i)
		}
	}
}

func TestSetIntersection(t *testing.T) {
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
		if !si.Contains(i) {
			t.Errorf("element not found: %d", i)
		}
	}
}

func TestSetSymmetricDifference(t *testing.T) {
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

	if len(ssd) != 10 {
		t.Errorf("union of size %d; expected %d", len(ssd), 3+4+3)
	}

	for i := 0; i < 3; i++ {
		if !ssd.Contains(i) {
			t.Errorf("element not found: %d", i)
		}
	}

	for i := 6; i < 10; i++ {
		if !ssd.Contains(i) {
			t.Errorf("element not found: %d", i)
		}
	}

	for i := 13; i < 16; i++ {
		if !ssd.Contains(i) {
			t.Errorf("element not found: %d", i)
		}
	}
}
