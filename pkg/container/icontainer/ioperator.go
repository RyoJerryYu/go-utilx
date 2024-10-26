package icontainer

type IOperator[T any] interface {
	Union(a, b T) T
	Subtract(a, b T) T
	Intersect(a, b T) T
	MergeAll(containers ...T) T // MergeAll do not mean to union all. It may not deduplicate.
	Equal(a, b T) bool
	Copy(a T) T
}
