package disjointx

import "github.com/RyoJerryYu/go-utilx/pkg/container/slicex"

// DisjointSetCore implements a disjoint-set data structure (also known as union-find).
// It maintains a collection of disjoint sets and supports efficient union and find operations.
// The implementation uses path compression and union by size for optimal performance.
type DisjointSetCore struct {
	// parents[idx] < 0 means the elements[idx] is a root of a set.
	// when parents[idx] > 0 (elements[idx] is not a root), parents[idx] means the parent of element[idx].
	// when parents[idx] < 0 (elements[idx] is a root), -parents[idx] means the size of the set.
	// elements with same root are in the same set.
	parents []int
}

// NewDisjointSetCore creates a new DisjointSetCore with n elements.
// Each element starts in its own singleton set.
func NewDisjointSetCore(n int) DisjointSetCore {
	parents := make([]int, n)
	for i := range parents {
		parents[i] = -1
	}
	return DisjointSetCore{
		parents: parents,
	}
}

// Find returns the root element of the set containing x.
// Uses path compression to maintain a flat tree structure.
// Time complexity: O(α(n)) amortized, where α is the inverse Ackermann function.
func (d *DisjointSetCore) Find(x int) int {
	if d.parents[x] < 0 {
		return x
	}
	d.parents[x] = d.Find(d.parents[x])
	return d.parents[x]
}

// Union merges the sets containing x and y.
// If x and y are already in the same set, nothing happens.
// Uses union by size to keep trees balanced.
// Time complexity: O(α(n)) amortized.
func (d *DisjointSetCore) Union(x, y int) {
	xRoot := d.Find(x)
	yRoot := d.Find(y)
	if xRoot == yRoot {
		return
	}

	if d.parents[xRoot] < d.parents[yRoot] {
		d.parents[xRoot] += d.parents[yRoot]
		d.parents[yRoot] = xRoot
	} else {
		d.parents[yRoot] += d.parents[xRoot]
		d.parents[xRoot] = yRoot
	}
}

// OrderedUnion merges the sets containing x and y, ensuring x's root becomes the new root.
// This operation is slower than Union but guarantees the root element.
// Example:
//
//	d := NewDisjointSetCore(3)
//	d.OrderedUnion(0, 1) // 0 becomes root
//	d.OrderedUnion(0, 2) // 0 remains root
//	d.Find(1) // returns 0
//	d.Find(2) // returns 0
func (d *DisjointSetCore) OrderedUnion(x, y int) {
	xRoot := d.Find(x)
	yRoot := d.Find(y)
	if xRoot == yRoot {
		return
	}

	d.parents[xRoot] += d.parents[yRoot]
	d.parents[yRoot] = xRoot
}

// SizeOf returns the size of the set containing x.
// Returns the number of elements in x's set.
func (d *DisjointSetCore) SizeOf(x int) int {
	return -d.parents[d.Find(x)]
}

// InSame returns true if x and y are in the same set.
func (d *DisjointSetCore) InSame(x, y int) bool {
	return d.Find(x) == d.Find(y)
}

// Roots returns all root elements in the disjoint sets.
// A root element is the representative element of its set.
func (d *DisjointSetCore) Roots() []int {
	res := make([]int, 0)
	for i, p := range d.parents {
		if p < 0 {
			res = append(res, i)
		}
	}
	return res
}

// CountGroups returns the number of disjoint sets.
func (d *DisjointSetCore) CountGroups() int {
	return len(d.Roots())
}

// Members returns all elements in the same set as x.
// The returned slice includes the root element.
func (d *DisjointSetCore) Members(x int) []int {
	root := d.Find(x)
	res := make([]int, 0)
	for i := range d.parents {
		if d.Find(i) == root {
			res = append(res, i)
		}
	}
	return res
}

// MembersWithoutRoot returns all elements in the same set as x, excluding the root element.
// Example:
//
//	d := NewDisjointSetCore(3)
//	d.Union(0, 1)
//	d.Union(0, 2)
//	d.MembersWithoutRoot(0) // returns [1, 2]
func (d *DisjointSetCore) MembersWithoutRoot(x int) []int {
	root := d.Find(x)
	res := make([]int, 0)
	for i := range d.parents {
		if d.Find(i) == root && i != root {
			res = append(res, i)
		}
	}
	return res
}

// MembersMap returns a map from root elements to their set members.
// The returned map's keys are root elements, and values are slices containing all elements in that set.
func (d *DisjointSetCore) MembersMap() map[int][]int {
	res := make(map[int][]int)
	for i := range d.parents {
		root := d.Find(i)
		res[root] = append(res[root], i)
	}
	return res
}

// MembersMapWithoutRoot returns a map from root elements to their non-root set members.
// Similar to MembersMap, but each set's root element is excluded from the member slice.
func (d *DisjointSetCore) MembersMapWithoutRoot() map[int][]int {
	res := make(map[int][]int)
	for i := range d.parents {
		root := d.Find(i)
		if root == i {
			continue
		}
		res[root] = append(res[root], i)
	}
	return res
}

// DisjointSet wraps DisjointSetCore to work with comparable types instead of just integers.
// Type parameter T must be comparable to support map operations.
type DisjointSet[T comparable] struct {
	core          DisjointSetCore
	elementIdxMap map[T]int // map<element, elementIndex>
	elements      []T       // map<elementIndex, element>
}

// NewDisjointSet creates a new DisjointSet with the given elements.
// Each element starts in its own singleton set.
// Example:
//
//	d := NewDisjointSet("a", "b", "c")
//	d.Union("a", "b")
//	d.Find("b") // returns "a", true
func NewDisjointSet[T comparable](elements ...T) DisjointSet[T] {
	core := NewDisjointSetCore(len(elements))
	idxMap := make(map[T]int)
	for idx, element := range elements {
		idxMap[element] = idx
	}
	return DisjointSet[T]{
		core:          core,
		elementIdxMap: idxMap,
		elements:      elements,
	}
}

func (d *DisjointSet[T]) idxToElement(i int) T {
	return d.elements[i]
}

func (d *DisjointSet[T]) have(x T) bool {
	_, ok := d.elementIdxMap[x]
	return ok
}

// Find returns the root element of x's set and true if x exists in the set.
// If x doesn't exist, returns x and false.
func (d *DisjointSet[T]) Find(x T) (T, bool) {
	if !d.have(x) {
		return x, false
	}
	rootElementIdx := d.core.Find(d.elementIdxMap[x])
	return d.elements[rootElementIdx], true
}

// Union merges the sets containing x and y.
// If either x or y doesn't exist in the set, nothing happens.
func (d *DisjointSet[T]) Union(x, y T) {
	if okx, oky := d.have(x), d.have(y); !okx || !oky {
		return
	}
	d.core.Union(d.elementIdxMap[x], d.elementIdxMap[y])
}

// OrderedUnion merges the sets containing x and y, ensuring x's value becomes the root.
// If either x or y doesn't exist in the set, nothing happens.
// Example:
//
//	d := NewDisjointSet("a", "b", "c")
//	d.OrderedUnion("a", "b") // "a" becomes root
//	d.OrderedUnion("a", "c") // "a" remains root
//	d.Find("b") // returns "a", true
func (d *DisjointSet[T]) OrderedUnion(x, y T) {
	if okx, oky := d.have(x), d.have(y); !okx || !oky {
		return
	}
	d.core.OrderedUnion(d.elementIdxMap[x], d.elementIdxMap[y])
}

func (d *DisjointSet[T]) SizeOf(x T) int {
	if !d.have(x) {
		return 0
	}
	return d.core.SizeOf(d.elementIdxMap[x])
}

func (d *DisjointSet[T]) InSame(x, y T) bool {
	if okx, oky := d.have(x), d.have(y); !okx || !oky {
		return false
	}
	return d.core.InSame(d.elementIdxMap[x], d.elementIdxMap[y])
}

// return root element of each groups
func (d *DisjointSet[T]) Roots() []T {
	rootElementIdxs := d.core.Roots()
	return slicex.To(rootElementIdxs, d.idxToElement)
}

func (d *DisjointSet[T]) CountGroups() int {
	return d.core.CountGroups()
}

// return elements with same group as x
func (d *DisjointSet[T]) Members(x T) []T {
	if !d.have(x) {
		return nil
	}
	memberIdxs := d.core.Members(d.elementIdxMap[x])
	return slicex.To(memberIdxs, d.idxToElement)
}

func (d *DisjointSet[T]) MembersWithoutRoot(x T) []T {
	if !d.have(x) {
		return nil
	}
	membersIdxs := d.core.MembersWithoutRoot(d.elementIdxMap[x])
	return slicex.To(membersIdxs, d.idxToElement)
}

// return elements array map by there root element.
func (d *DisjointSet[T]) MembersMap() map[T][]T {
	membersIndexMap := d.core.MembersMap()
	res := make(map[T][]T)
	for root, memberIdxs := range membersIndexMap {
		res[d.elements[root]] = slicex.To(memberIdxs, d.idxToElement)
	}
	return res
}

// MembersMapWithoutRoot returns a map from root elements to their non-root set members.
// Similar to MembersMap, but each set's root element is excluded from the member slice.
// Example:
//
//	d := NewDisjointSet("a", "b", "c")
//	d.Union("a", "b")
//	d.MembersMapWithoutRoot() // returns {"a": ["b"], "c": []}
func (d *DisjointSet[T]) MembersMapWithoutRoot() map[T][]T {
	membersIndexMap := d.core.MembersMapWithoutRoot()
	res := make(map[T][]T)
	for root, memberIdxs := range membersIndexMap {
		res[d.elements[root]] = slicex.To(memberIdxs, d.idxToElement)
	}
	return res
}
