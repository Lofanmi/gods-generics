package contracts

// Comparator will make type assertion (see IntComparator for example),
// which will panic if a or b are not of the asserted type.
//
// Should return a number:
//
//	negative , if a < b
//	zero     , if a == b
//	positive , if a > b
type Comparator[T comparable] func(a, b T) int
