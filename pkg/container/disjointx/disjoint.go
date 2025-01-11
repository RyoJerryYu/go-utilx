package disjointx

import "github.com/RyoJerryYu/go-utilx/pkg/container/slicex"

type DisjointSetCore struct {
	// parents[idx] < 0 means the elements[idx] is a root of a set.
	// when parents[idx] > 0 (elements[idx] is not a root), parents[idx] means the parent of element[idx].
	// when parents[idx] < 0 (elements[idx] is a root), -parents[idx] means the size of the set.
	// elements with same root are in the same set.
	parents []int
}

func NewDisjointSetCore(n int) DisjointSetCore {
	parents := make([]int, n)
	for i := range parents {
		parents[i] = -1
	}
	return DisjointSetCore{
		parents: parents,
	}
}

func (d *DisjointSetCore) Find(x int) int {
	if d.parents[x] < 0 {
		return x
	}
	d.parents[x] = d.Find(d.parents[x])
	return d.parents[x]
}

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

// 将 y 集合并入 x 集合
// 固定最后会以 x 集合的根为两个集合的根
// 均摊时间复杂度上会比 Union 慢
//
// Join y set into x set
// The root of the x set will be root of the two sets
// The amortized time complexity will be slower than Union
func (d *DisjointSetCore) OrderedUnion(x, y int) {
	xRoot := d.Find(x)
	yRoot := d.Find(y)
	if xRoot == yRoot {
		return
	}

	d.parents[xRoot] += d.parents[yRoot]
	d.parents[yRoot] = xRoot
}

func (d *DisjointSetCore) SizeOf(x int) int {
	return -d.parents[d.Find(x)]
}

func (d *DisjointSetCore) InSame(x, y int) bool {
	return d.Find(x) == d.Find(y)
}

func (d *DisjointSetCore) Roots() []int {
	res := make([]int, 0)
	for i, p := range d.parents {
		if p < 0 {
			res = append(res, i)
		}
	}
	return res
}

func (d *DisjointSetCore) CountGroups() int {
	return len(d.Roots())
}

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

func (d *DisjointSetCore) MembersMap() map[int][]int {
	res := make(map[int][]int)
	for i := range d.parents {
		root := d.Find(i)
		res[root] = append(res[root], i)
	}
	return res
}

// MembersMapWithoutRoot return members index array map by there root index.
// The root of each set is not included in the array.
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

type DisjointSet[T comparable] struct {
	core          DisjointSetCore
	elementIdxMap map[T]int // map<element, elementIndex>
	elements      []T       // map<elementIndex, element>
}

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

func (d *DisjointSet[T]) Find(x T) (T, bool) {
	if !d.have(x) {
		return x, false
	}
	rootElementIdx := d.core.Find(d.elementIdxMap[x])
	return d.elements[rootElementIdx], true
}

func (d *DisjointSet[T]) Union(x, y T) {
	if okx, oky := d.have(x), d.have(y); !okx || !oky {
		return
	}
	d.core.Union(d.elementIdxMap[x], d.elementIdxMap[y])
}

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

// return elements array without root map by there root element.
func (d *DisjointSet[T]) MembersMapWithoutRoot() map[T][]T {
	membersIndexMap := d.core.MembersMapWithoutRoot()
	res := make(map[T][]T)
	for root, memberIdxs := range membersIndexMap {
		res[d.elements[root]] = slicex.To(memberIdxs, d.idxToElement)
	}
	return res
}
