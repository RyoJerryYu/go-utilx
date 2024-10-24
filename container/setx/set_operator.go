package setx

import "github.com/RyoJerryYu/go-utilx/container/icontainer"

type SetOperator[T comparable] struct{}

func (SetOperator[T]) Union(a, b Set[T]) Set[T]                 { return a.Union(b) }
func (SetOperator[T]) Subtract(a, b Set[T]) Set[T]              { return a.Subtract(b) }
func (SetOperator[T]) Intersect(a, b Set[T]) Set[T]             { return a.Intersect(b) }
func (SetOperator[T]) Merge(a Set[T], others ...Set[T]) Set[T]  { return a.MergeSet(others...) }
func (SetOperator[T]) Remove(a Set[T], others ...Set[T]) Set[T] { return a.RemoveSet(others...) }
func (SetOperator[T]) Equal(a Set[T], b Set[T]) bool            { return a.Equal(b) }
func (SetOperator[T]) Copy(a Set[T]) Set[T]                     { return a.Copy() }

var _ icontainer.IOperator[Set[string]] = SetOperator[string]{}
