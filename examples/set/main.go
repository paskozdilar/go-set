package main

import (
	"fmt"
	"slices"

	set "github.com/paskozdilar/go-set"
)

func main() {
	m := make(set.Set[int])

	m.Insert(420)
	fmt.Println(m.Contains(420)) // true
	m.Remove(420)

	m1 := make(set.Set[int])
	m2 := make(set.Set[int])

	for i := 0; i < 6; i++ {
		m1.Insert(i)
	}
	for i := 3; i < 9; i++ {
		m2.Insert(i)
	}

	mu := m1.Union(m2)
	eu := set.Elements(mu) // slice containing elements of set
	slices.Sort(eu)
	// 0, 1, 2, 3, 4, 5, 6, 7, 8
	for _, e := range eu {
		fmt.Println(e)
	}

	mi := m1.Intersection(m2)
	ei := set.Elements(mi)
	slices.Sort(ei)
	// 3, 4, 5
	for _, e := range ei {
		fmt.Println(e)
	}

	msd := m1.SymmetricDifference(m2)
	esd := set.Elements(msd)
	slices.Sort(esd)
	// 0, 1, 2, 6, 7, 8
	for _, e := range esd {
		fmt.Println(e)
	}
}
