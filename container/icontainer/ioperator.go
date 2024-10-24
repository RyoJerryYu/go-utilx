package icontainer

type IOperator[T any] interface {
	Union(a, b T) T
	Subtract(a, b T) T
	Intersect(a, b T) T
	Merge(a T, others ...T) T
	Remove(a T, others ...T) T
	Equal(a, b T) bool
	Copy(a T) T
}
