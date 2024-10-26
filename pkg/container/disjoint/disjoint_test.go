package disjoint

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDisjointSet(t *testing.T) {
	elements := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	// length is same
	d := NewDisjointSet(len(elements))
	assert.Equal(t, len(elements), len(d.parents))

	// all elements are roots
	roots := d.Roots()
	assert.Equal(t, len(elements), len(roots))
	assert.Equal(t, len(elements), d.CountGroups())
	assert.False(t, d.InSame(0, 1))
	assert.False(t, d.InSame(0, 2))

	// union one
	d.Union(0, 1)
	roots = d.Roots()
	assert.Equal(t, len(elements)-1, len(roots))
	assert.Equal(t, len(elements)-1, d.CountGroups())
	assert.True(t, d.InSame(0, 1))
	assert.False(t, d.InSame(0, 2))

	// union more
	d.Union(0, 2)
	d.Union(0, 3)
	d.Union(4, 5)
	d.Union(6, 7)
	d.Union(8, 9)
	d.Union(0, 4)

	// 0: 1,2,3,4,5
	// 6: 7
	// 8: 9
	roots = d.Roots()
	assert.Equal(t, 3, len(roots))
	assert.Equal(t, 3, d.CountGroups())
	assert.True(t, d.InSame(0, 1))
	assert.True(t, d.InSame(1, 4))
	assert.True(t, d.InSame(2, 5))
	assert.True(t, d.InSame(6, 7))
	assert.False(t, d.InSame(0, 7))
}

// test for OrderedUnion, MembersWithoutRoot
func TestDisjointSet_Ordered(t *testing.T) {
	elements := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	d1 := NewDisjointSet(len(elements))

	d1.OrderedUnion(0, 1)
	d1.OrderedUnion(0, 2)
	d1.OrderedUnion(0, 3)
	d1.OrderedUnion(4, 5)
	d1.OrderedUnion(6, 7)
	d1.OrderedUnion(8, 9)
	d1.OrderedUnion(0, 4)
	// 0: 1,2,3,4,5
	// 6: 7
	// 8: 9
	roots := d1.Roots()
	assert.Equal(t, 3, len(roots))
	assert.Equal(t, 3, d1.CountGroups())
	assert.Contains(t, roots, 0)
	assert.Contains(t, roots, 6)
	assert.Contains(t, roots, 8)

	membersMap := d1.MembersMapWithoutRoot()
	assert.Len(t, membersMap, 3)
	assert.Len(t, membersMap[0], 5)
	assert.Len(t, membersMap[6], 1)
	assert.Len(t, membersMap[8], 1)

	d2 := NewDisjointSet(len(elements))

	// change the order of union don't affect the root result
	d2.OrderedUnion(4, 5)
	d2.OrderedUnion(0, 1)
	d2.OrderedUnion(6, 7)
	d2.OrderedUnion(0, 4)
	d2.OrderedUnion(0, 2)
	d2.OrderedUnion(8, 9)
	d2.OrderedUnion(0, 3)

	roots = d1.Roots()
	assert.Equal(t, 3, len(roots))
	assert.Equal(t, 3, d1.CountGroups())
	assert.Contains(t, roots, 0)
	assert.Contains(t, roots, 6)
	assert.Contains(t, roots, 8)
}

func TestDisjointSetElement(t *testing.T) {
	elements := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	d := NewDisjointSetElement(elements...)

	// all elements are roots
	roots := d.Roots()
	assert.Equal(t, len(elements), len(roots))
	assert.Equal(t, len(elements), d.CountGroups())
	assert.False(t, d.InSame("0", "1"))
	assert.False(t, d.InSame("0", "2"))

	// union one
	d.Union("0", "1")
	roots = d.Roots()
	assert.Equal(t, len(elements)-1, len(roots))
	assert.Equal(t, len(elements)-1, d.CountGroups())
	assert.True(t, d.InSame("0", "1"))
	assert.False(t, d.InSame("0", "2"))

	// union more
	d.Union("0", "2")
	d.Union("0", "3")
	d.Union("4", "5")
	d.Union("6", "7")
	d.Union("8", "9")
	d.Union("0", "4")

	// 0: 1,2,3,4,5
	// 6: 7
	// 8: 9
	roots = d.Roots()
	assert.Equal(t, 3, len(roots))
	assert.Equal(t, 3, d.CountGroups())
	assert.True(t, d.InSame("0", "1"))
	assert.True(t, d.InSame("1", "4"))
	assert.True(t, d.InSame("2", "5"))
	assert.True(t, d.InSame("6", "7"))
	assert.False(t, d.InSame("0", "7"))
}

func TestDisjointSetElement_Ordered(t *testing.T) {
	elements := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	d1 := NewDisjointSetElement(elements...)

	d1.OrderedUnion("0", "1")
	d1.OrderedUnion("0", "2")
	d1.OrderedUnion("0", "3")
	d1.OrderedUnion("4", "5")
	d1.OrderedUnion("6", "7")
	d1.OrderedUnion("8", "9")
	d1.OrderedUnion("0", "4")
	// 0: 1,2,3,4,5
	// 6: 7
	// 8: 9
	roots := d1.Roots()
	assert.Equal(t, 3, len(roots))
	assert.Equal(t, 3, d1.CountGroups())
	assert.Contains(t, roots, "0")
	assert.Contains(t, roots, "6")
	assert.Contains(t, roots, "8")

	membersMap := d1.MembersMapWithoutRoot()
	assert.Len(t, membersMap, 3)
	assert.Len(t, membersMap["0"], 5)
	assert.Len(t, membersMap["6"], 1)
	assert.Len(t, membersMap["8"], 1)

	d2 := NewDisjointSetElement(elements...)

	// change the order of union don't affect the root result
	d2.OrderedUnion("4", "5")
	d2.OrderedUnion("0", "1")
	d2.OrderedUnion("6", "7")
	d2.OrderedUnion("0", "4")
	d2.OrderedUnion("0", "2")
	d2.OrderedUnion("8", "9")
	d2.OrderedUnion("0", "3")

	roots = d2.Roots()
	assert.Equal(t, 3, len(roots))
	assert.Equal(t, 3, d1.CountGroups())
	assert.Contains(t, roots, "0")
	assert.Contains(t, roots, "6")
	assert.Contains(t, roots, "8")
}
