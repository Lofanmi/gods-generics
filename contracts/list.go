package contracts

// List interface that all lists implement
type List[T comparable] interface {
	Get(index int) (T, bool)
	Remove(index int)
	Add(values ...T)
	Contains(values ...T) bool
	Sort(comparator Comparator[T])
	Swap(index1, index2 int)
	Insert(index int, values ...T)
	Set(index int, value T)

	Container[T]
	// Empty() bool
	// Size() int
	// Clear()
	// Values() []T
	// String() string
}
