// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package utils

import (
	"sort"

	"github.com/Lofanmi/gods-generics/contracts"
)

// Sort sorts values (in-place) with respect to the given comparator.
//
// Uses Go's sort (hybrid of quicksort for large and then insertion sort for smaller slices).
func Sort[T comparable](values []T, comparator contracts.Comparator[T]) {
	sort.Sort(sortable[T]{values, comparator})
}

type sortable[T comparable] struct {
	values     []T
	comparator contracts.Comparator[T]
}

func (s sortable[T]) Len() int {
	return len(s.values)
}

func (s sortable[T]) Swap(i, j int) {
	s.values[i], s.values[j] = s.values[j], s.values[i]
}

func (s sortable[T]) Less(i, j int) bool {
	return s.comparator(s.values[i], s.values[j]) < 0
}
