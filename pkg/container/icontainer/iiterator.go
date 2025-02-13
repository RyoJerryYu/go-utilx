package icontainer

// IIterator represents an iterator interface for sequential access to container elements.
type IIterator[T any] interface {
	// First returns the first element and true if the iterator is not empty.
	// If the iterator is empty, returns the zero value and false.
	First() (T, bool)
	// Next returns the next element and true if there are more elements.
	// If there are no more elements, returns the zero value and false.
	Next() (T, bool)
}

// IterContainer represents a container that can provide an iterator.
type IterContainer[T any] interface {
	// ToIter returns an iterator for the container.
	ToIter() IIterator[T]
}

// ForEach executes the given function for each element in the iterator.
func ForEach[T any](iterable IIterator[T], fn func(T)) {
	for item, ok := iterable.First(); ok; item, ok = iterable.Next() {
		fn(item)
	}
}

// Collect converts an iterable container into a slice.
func Collect[T any](container IterContainer[T]) []T {
	var result []T
	ForEach(container.ToIter(), func(item T) {
		result = append(result, item)
	})
	return result
}

// MapBy creates a map from an iterable container using the provided function
// to generate keys for each element.
func MapBy[T any, I comparable](container IterContainer[T], fn func(T) I) map[I]T {
	out := make(map[I]T)
	ForEach(container.ToIter(), func(v T) {
		out[fn(v)] = v
	})
	return out
}

// GroupBy groups elements from an iterable container into a map of slices
// using the provided function to generate keys for each element.
func GroupBy[T any, I comparable](container IterContainer[T], fn func(T) I) map[I][]T {
	out := make(map[I][]T)
	ForEach(container.ToIter(), func(v T) {
		k := fn(v)
		out[k] = append(out[k], v)
	})
	return out
}
