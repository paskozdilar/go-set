package main

import (
	"fmt"
	"slices"

	set "github.com/paskozdilar/go-set"
)

func main() {
	m := make(map[int]string)

	set.Insert(m, 420)
	fmt.Println(set.Contains(m, 420)) // true
	set.Remove(m, 420)

	m1 := make(map[int]string)
	m2 := make(map[int]string)

	for i := 0; i < 6; i++ {
		set.Insert(m1, i)
	}
	for i := 3; i < 9; i++ {
		set.Insert(m2, i)
	}

	mu := set.Union(m1, m2)
	eu := set.Elements(mu) // slice containing elements of set
	slices.Sort(eu)
	// 0, 1, 2, 3, 4, 5, 6, 7, 8
	for _, e := range eu {
		fmt.Println(e)
	}

	mi := set.Intersection(m1, m2)
	ei := set.Elements(mi)
	slices.Sort(ei)
	// 3, 4, 5
	for _, e := range ei {
		fmt.Println(e)
	}

	msd := set.SymmetricDifference(m1, m2)
	esd := set.Elements(msd)
	slices.Sort(esd)
	// 0, 1, 2, 6, 7, 8
	for _, e := range esd {
		fmt.Println(e)
	}
}
