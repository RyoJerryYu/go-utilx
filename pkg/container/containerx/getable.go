package containerx

import (
	"cmp"
	"slices"

	"github.com/RyoJerryYu/go-utilx/pkg/container/slicex"
)

type IdGetable[ID comparable] interface{ GetId() ID }
type NameGetable[Name comparable] interface{ GetName() Name }
type PiddGetable[Pid comparable] interface{ GetPid() Pid }

func ToIds[ID comparable, T IdGetable[ID]](s []T) []ID           { return slicex.To(s, T.GetId) }
func ToNames[Name comparable, T NameGetable[Name]](s []T) []Name { return slicex.To(s, T.GetName) }
func ToPids[PID comparable, T PiddGetable[PID]](s []T) []PID     { return slicex.To(s, T.GetPid) }

func MapByIds[ID comparable, T IdGetable[ID]](s []T) map[ID]T { return slicex.MapBy(s, T.GetId) }
func MapByNames[Name comparable, T NameGetable[Name]](s []T) map[Name]T {
	return slicex.MapBy(s, T.GetName)
}

func GroupByPids[PID comparable, T PiddGetable[PID]](s []T) map[PID][]T {
	return slicex.GroupBy(s, T.GetPid)
}
func GroupByNames[Name comparable, T NameGetable[Name]](s []T) map[Name][]T {
	return slicex.GroupBy(s, T.GetName)
}

func ChunkBy[T any, By cmp.Ordered](s []T, by func(T) By) [][]T {
	groupBy := slicex.GroupBy(s, by)
	keys := slicex.FromKey(groupBy)
	slices.Sort(keys)
	return slicex.To(keys, func(key By) []T { return groupBy[key] })
}
func ChunkByIds[ID cmp.Ordered, T IdGetable[ID]](s []T) [][]T         { return ChunkBy(s, T.GetId) }
func ChunkByNames[Name cmp.Ordered, T NameGetable[Name]](s []T) [][]T { return ChunkBy(s, T.GetName) }
func ChunkByPids[PID cmp.Ordered, T PiddGetable[PID]](s []T) [][]T    { return ChunkBy(s, T.GetPid) }
