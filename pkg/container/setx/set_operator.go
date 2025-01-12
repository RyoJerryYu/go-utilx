package setx

import (
	"github.com/RyoJerryYu/go-utilx/pkg/container/icontainer"
)

type Operator[T comparable] struct{}

func (Operator[T]) Union(a, b Set[T]) Set[T]       { return a.Union(b) }
func (Operator[T]) Subtract(a, b Set[T]) Set[T]    { return a.Subtract(b) }
func (Operator[T]) Intersect(a, b Set[T]) Set[T]   { return a.Intersect(b) }
func (Operator[T]) MergeAll(sets ...Set[T]) Set[T] { return SetMergeAll(sets...) }
func (Operator[T]) Equal(a Set[T], b Set[T]) bool  { return a.Equal(b) }
func (Operator[T]) Copy(a Set[T]) Set[T]           { return a.Copy() }

var _ icontainer.IOperator[Set[string]] = Operator[string]{}
