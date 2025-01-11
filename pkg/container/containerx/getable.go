package containerx

import "github.com/RyoJerryYu/go-utilx/pkg/container/slicex"

type IdGetable[ID comparable] interface{ GetId() ID }
type NameGetable[Name comparable] interface{ GetName() Name }

func ToIds[ID comparable, T IdGetable[ID]](s []T) []ID           { return slicex.To(s, T.GetId) }
func ToNames[Name comparable, T NameGetable[Name]](s []T) []Name { return slicex.To(s, T.GetName) }

func MapByIds[ID comparable, T IdGetable[ID]](s []T) map[ID]T { return slicex.MapBy(s, T.GetId) }
func MapByNames[Name comparable, T NameGetable[Name]](s []T) map[Name]T {
	return slicex.MapBy(s, T.GetName)
}
