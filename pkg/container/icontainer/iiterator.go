package icontainer

type IIterator[T any] interface {
	First() (T, bool)
	Next() (T, bool)
}

type IterContainer[T any] interface {
	ToIter() IIterator[T]
}

func ForEach[T any](iterable IIterator[T], fn func(T)) {
	for item, ok := iterable.First(); ok; item, ok = iterable.Next() {
		fn(item)
	}
}

func Collect[T any](container IterContainer[T]) []T {
	var result []T
	ForEach(container.ToIter(), func(item T) {
		result = append(result, item)
	})
	return result
}

func MapBy[T any, I comparable](container IterContainer[T], fn func(T) I) map[I]T {
	out := make(map[I]T)
	ForEach(container.ToIter(), func(v T) {
		out[fn(v)] = v
	})
	return out
}

func GroupBy[T any, I comparable](container IterContainer[T], fn func(T) I) map[I][]T {
	out := make(map[I][]T)
	ForEach(container.ToIter(), func(v T) {
		k := fn(v)
		out[k] = append(out[k], v)
	})
	return out
}
