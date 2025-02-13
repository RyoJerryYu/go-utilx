package icontainer

// IOperator defines operations that can be performed between two containers of the same type.
type IOperator[T any] interface {
	// Union returns a new container containing all elements from both containers.
	Union(a, b T) T
	// Subtract returns a new container containing elements from a that are not in b.
	Subtract(a, b T) T
	// Intersect returns a new container containing elements that exist in both containers.
	Intersect(a, b T) T
	// MergeAll combines multiple containers into a single container.
	// Note: This operation may not deduplicate elements.
	MergeAll(containers ...T) T
	// Equal returns true if both containers contain the same elements.
	Equal(a, b T) bool
	// Copy returns a deep copy of the given container.
	Copy(a T) T
}
