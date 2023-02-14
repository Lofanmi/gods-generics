package contracts

// Container is base interface that all data structures implement.
type Container[T comparable] interface {
	Empty() bool
	Size() int
	Clear()
	Values() []T
	String() string
}
