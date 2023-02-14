package contracts

import "golang.org/x/exp/constraints"

type Map[K constraints.Ordered, V comparable] interface {
	Len() int
	Keys() []K
	Values() []V
	Range(func(K, V) bool)
	Get(K) V
	GetExist(K) (V, bool)
	Set(K, V)
	Del(K)
}
