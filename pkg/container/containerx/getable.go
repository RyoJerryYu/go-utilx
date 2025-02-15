package containerx

import (
	"cmp"
	"slices"

	"github.com/RyoJerryYu/go-utilx/pkg/container/slicex"
)

// IdGetable represents a type that can provide an ID of type ID.
// Type parameter ID must be comparable.
type IdGetable[ID comparable] interface{ GetId() ID }

// NameGetable represents a type that can provide a Name of type Name.
// Type parameter Name must be comparable.
type NameGetable[Name comparable] interface{ GetName() Name }

// PiddGetable represents a type that can provide a Parent ID (Pid) of type Pid.
// Type parameter Pid must be comparable.
type PiddGetable[Pid comparable] interface{ GetPid() Pid }

// ToIds extracts IDs from a slice of IdGetable elements.
// The order of IDs in the result matches the order of elements in the input slice.
// Example:
//
//	type User struct{ id int }
//	func (u User) GetId() int { return u.id }
//	users := []User{{1}, {2}, {3}}
//	ids := ToIds(users)
//	// ids = []int{1, 2, 3}
func ToIds[ID comparable, T IdGetable[ID]](s []T) []ID { return slicex.To(s, T.GetId) }

// ToNames extracts names from a slice of NameGetable elements.
// The order of names in the result matches the order of elements in the input slice.
// Example:
//
//	type User struct{ name string }
//	func (u User) GetName() string { return u.name }
//	users := []User{{"Alice"}, {"Bob"}, {"Charlie"}}
//	names := ToNames(users)
//	// names = []string{"Alice", "Bob", "Charlie"}
func ToNames[Name comparable, T NameGetable[Name]](s []T) []Name { return slicex.To(s, T.GetName) }

// ToPids extracts parent IDs from a slice of PiddGetable elements.
// The order of PIDs in the result matches the order of elements in the input slice.
// Example:
//
//	type Node struct{ pid int }
//	func (n Node) GetPid() int { return n.pid }
//	nodes := []Node{{1}, {2}, {2}}
//	pids := ToPids(nodes)
//	// pids = []int{1, 2, 2}
func ToPids[PID comparable, T PiddGetable[PID]](s []T) []PID { return slicex.To(s, T.GetPid) }

// MapByIds creates a map from a slice of IdGetable elements using their IDs as keys.
// If multiple elements have the same ID, the last one will be kept.
// Example:
//
//	type User struct{ id int }
//	func (u User) GetId() int { return u.id }
//	users := []User{{1}, {2}, {3}}
//	userMap := MapByIds(users)
//	// userMap = map[int]User{1: {1}, 2: {2}, 3: {3}}
func MapByIds[ID comparable, T IdGetable[ID]](s []T) map[ID]T { return slicex.MapBy(s, T.GetId) }

// MapByNames creates a map from a slice of NameGetable elements using their names as keys.
// If multiple elements have the same name, the last one will be kept.
// Example:
//
//	type User struct{ name string }
//	func (u User) GetName() string { return u.name }
//	users := []User{{"Alice"}, {"Bob"}, {"Charlie"}}
//	userMap := MapByNames(users)
//	// userMap = map[string]User{"Alice": {"Alice"}, "Bob": {"Bob"}, "Charlie": {"Charlie"}}
func MapByNames[Name comparable, T NameGetable[Name]](s []T) map[Name]T {
	return slicex.MapBy(s, T.GetName)
}

// FilterBy filters a slice of elements based on a set of values.
// The filter function returns true if the element should be included in the result.
// Example:
//
//	type User struct{ id int }
//	func (u User) GetId() int { return u.id }
//	users := []User{{1}, {2}, {3}}
//	filtered := FilterBy(users, User.GetId, 1, 3)
//	// filtered = []User{{1}, {3}}
func FilterBy[T any, By comparable](s []T, by func(T) By, values ...By) []T {
	valueSet := slicex.ToSet(values)
	return slicex.Filter(s, func(t T) bool {
		_, ok := valueSet[by(t)]
		return ok
	})
}

// FilterByIds filters a slice of elements based on a set of IDs.
// The filter function returns true if the element should be included in the result.
// Example:
//
//	type User struct{ id int }
//	func (u User) GetId() int { return u.id }
//	users := []User{{1}, {2}, {3}}
//	filtered := FilterByIds(users, 1, 3)
//	// filtered = []User{{1}, {3}}
func FilterByIds[ID comparable, T IdGetable[ID]](s []T, ids ...ID) []T {
	return FilterBy(s, T.GetId, ids...)
}

// FilterByNames filters a slice of elements based on a set of names.
// The filter function returns true if the element should be included in the result.
// Example:
//
//	type User struct{ name string }
//	func (u User) GetName() string { return u.name }
//	users := []User{{"Alice"}, {"Bob"}, {"Charlie"}}
//	filtered := FilterByNames(users, "Alice", "Charlie")
//	// filtered = []User{{"Alice"}, {"Charlie"}}
func FilterByNames[Name comparable, T NameGetable[Name]](s []T, names ...Name) []T {
	return FilterBy(s, T.GetName, names...)
}

// FilterByPids filters a slice of elements based on a set of parent IDs.
// The filter function returns true if the element should be included in the result.
// Example:
//
//	type Node struct{ pid int }
//	func (n Node) GetPid() int { return n.pid }
//	nodes := []Node{{1}, {2}, {2}}
//	filtered := FilterByPids(nodes, 1, 2)
//	// filtered = []Node{{1}, {2}}
func FilterByPids[PID comparable, T PiddGetable[PID]](s []T, pids ...PID) []T {
	return FilterBy(s, T.GetPid, pids...)
}

// GroupByPids groups elements by their parent IDs.
// Elements with the same parent ID will be grouped together in a slice.
// Example:
//
//	type Node struct{ pid int }
//	func (n Node) GetPid() int { return n.pid }
//	nodes := []Node{{1}, {1}, {2}}
//	groups := GroupByPids(nodes)
//	// groups = map[int][]Node{1: {{1}, {1}}, 2: {{2}}}
func GroupByPids[PID comparable, T PiddGetable[PID]](s []T) map[PID][]T {
	return slicex.GroupBy(s, T.GetPid)
}

// GroupByNames groups elements by their names.
// Elements with the same name will be grouped together in a slice.
// Example:
//
//	type User struct{ name string }
//	func (u User) GetName() string { return u.name }
//	users := []User{{"Alice"}, {"Bob"}, {"Bob"}}
//	groups := GroupByNames(users)
//	// groups = map[string][]User{"Alice": {{"Alice"}}, "Bob": {{"Bob"}, {"Bob"}}}
func GroupByNames[Name comparable, T NameGetable[Name]](s []T) map[Name][]T {
	return slicex.GroupBy(s, T.GetName)
}

// ChunkBy splits a slice into chunks based on a key function.
// Elements are first grouped by the key function, then sorted by keys.
// The result is a slice of slices where each inner slice contains elements with the same key.
// Type parameter By must be ordered to support sorting.
// Example:
//
//	numbers := []int{1, 1, 2, 3, 3}
//	chunks := ChunkBy(numbers, func(n int) int { return n })
//	// chunks = [][]int{{1, 1}, {2}, {3, 3}}
func ChunkBy[T any, By cmp.Ordered](s []T, by func(T) By) [][]T {
	groupBy := slicex.GroupBy(s, by)
	keys := slicex.FromKey(groupBy)
	slices.Sort(keys)
	return slicex.To(keys, func(key By) []T { return groupBy[key] })
}

// ChunkByIds splits a slice into chunks based on element IDs.
// Elements with the same ID will be in the same chunk.
// Chunks are ordered by ID.
// Type parameter ID must be ordered to support sorting.
func ChunkByIds[ID cmp.Ordered, T IdGetable[ID]](s []T) [][]T { return ChunkBy(s, T.GetId) }

// ChunkByNames splits a slice into chunks based on element names.
// Elements with the same name will be in the same chunk.
// Chunks are ordered by name.
// Type parameter Name must be ordered to support sorting.
// Example:
//
//	type User struct{ name string }
//	func (u User) GetName() string { return u.name }
//	users := []User{{"Alice"}, {"Bob"}, {"Bob"}, {"Charlie"}}
//	chunks := ChunkByNames(users)
//	// chunks = [][]User{[{"Alice"}], [{"Bob"}, {"Bob"}], [{"Charlie"}]}
func ChunkByNames[Name cmp.Ordered, T NameGetable[Name]](s []T) [][]T { return ChunkBy(s, T.GetName) }

// ChunkByPids splits a slice into chunks based on parent IDs.
// Elements with the same PID will be in the same chunk.
// Chunks are ordered by PID.
// Type parameter PID must be ordered to support sorting.
// Example:
//
//	type Node struct{ pid int }
//	func (n Node) GetPid() int { return n.pid }
//	nodes := []Node{{1}, {2}, {2}, {3}}
//	chunks := ChunkByPids(nodes)
//	// chunks = [][]Node{[{1}], [{2}, {2}], [{3}]}
func ChunkByPids[PID cmp.Ordered, T PiddGetable[PID]](s []T) [][]T { return ChunkBy(s, T.GetPid) }
