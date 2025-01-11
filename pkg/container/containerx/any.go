package containerx

import "github.com/RyoJerryYu/go-utilx/pkg/container/slicex"

func Any[T any](in T) any {
	return in
}

func ToAny[T any](in []T) []any {
	return slicex.To(in, Any)
}
