package treex

import "github.com/RyoJerryYu/go-utilx/pkg/container/icontainer"

// TreeNode represents a node in a tree structure.
// Type parameters:
//   - ID: The type of node identifier, must be comparable
//   - T: The type of value stored in the node
type TreeNode[ID comparable, T any] interface {
	// GetId returns the unique identifier of this node
	GetId() ID
	// GetPath returns the path from root to this node, including this node's ID
	GetPath() []ID
	// GetValue returns the value stored in this node
	GetValue() T
	// Size returns the total number of nodes in the subtree rooted at this node
	Size() int
}

// NodeOperator defines operations that can be performed on tree nodes.
// Type parameters:
//   - ID: The type of node identifier, must be comparable
//   - T: The type of value stored in the node
//   - N: The concrete type implementing TreeNode interface
type NodeOperator[ID comparable, T any, N TreeNode[ID, T]] interface {
	// Parent returns the parent node of the given node
	// Returns nil if the node is a root
	Parent(node N) N
	// Children returns all direct child nodes of the given node
	Children(node N) []N
	// AddChild adds a child node to the given parent node
	// This operation should update the child's parent reference and path
	AddChild(node N, child N)
}

// PreorderTraversal performs a pre-order traversal of the tree.
// The traversal visits the root node first, then recursively traverses each subtree.
// Example:
//
//	tree:     1
//	        /   \
//	       2     3
//	      /
//	     4
//
//	PreorderTraversal will visit nodes in order: 1, 2, 4, 3
func PreorderTraversal[ID comparable, T any, N TreeNode[ID, T], O NodeOperator[ID, T, N]](roots []N, operator O, fn func(N)) {
	for _, root := range roots {
		fn(root)
		PreorderTraversal(operator.Children(root), operator, fn)
	}
}

// PostorderTraversal performs a post-order traversal of the tree.
// The traversal recursively visits all subtrees first, then visits the root node.
// Example:
//
//	tree:     1
//	        /   \
//	       2     3
//	      /
//	     4
//
//	PostorderTraversal will visit nodes in order: 4, 2, 3, 1
func PostorderTraversal[ID comparable, T any, N TreeNode[ID, T], O NodeOperator[ID, T, N]](roots []N, operator O, fn func(N)) {
	for _, root := range roots {
		PostorderTraversal(operator.Children(root), operator, fn)
		fn(root)
	}
}

// LevelOrderTraversal performs a level-order (breadth-first) traversal of the tree.
// The traversal visits all nodes at the current depth before moving to nodes at the next depth level.
// Example:
//
//	tree:     1
//	        /   \
//	       2     3
//	      /
//	     4
//
//	LevelOrderTraversal will visit nodes in order: 1, 2, 3, 4
func LevelOrderTraversal[ID comparable, T any, N TreeNode[ID, T], O NodeOperator[ID, T, N]](roots []N, operator O, fn func(N)) {
	queue := roots
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		fn(node)
		queue = append(queue, operator.Children(node)...)
	}
}

// Trees represents a forest (collection of trees).
// Type parameters:
//   - ID: The type of node identifier, must be comparable
//   - T: The type of value stored in nodes
//   - N: The concrete type implementing TreeNode interface
//   - O: The concrete type implementing NodeOperator interface
type Trees[ID comparable, T any, N TreeNode[ID, T], O NodeOperator[ID, T, N]] struct {
	Roots []N
}

// New creates a new empty Trees instance.
// Type parameters follow the same constraints as Trees.
func New[ID comparable, T any, N TreeNode[ID, T], O NodeOperator[ID, T, N]]() *Trees[ID, T, N, O] {
	return &Trees[ID, T, N, O]{}
}

// GetRoots returns the root nodes of all trees in the forest.
func (t *Trees[ID, T, N, O]) GetRoots() []N { return t.Roots }

// Len returns the total number of nodes in all trees.
// Implements Container.Len().
func (t *Trees[ID, T, N, O]) Len() int {
	cnt := 0
	for _, root := range t.Roots {
		cnt += root.Size()
	}
	return cnt
}

// IsEmpty returns true if there are no trees in the forest.
// Implements Container.IsEmpty().
func (t *Trees[ID, T, N, O]) IsEmpty() bool {
	return len(t.Roots) == 0
}

// Clear removes all trees from the forest.
// Implements Container.Clear().
func (t *Trees[ID, T, N, O]) Clear() {
	t.Roots = []N{}
}

func (t *Trees[ID, T, N, O]) PreorderTraversal(fn func(N)) {
	PreorderTraversal(t.Roots, *new(O), fn)
}

func (t *Trees[ID, T, N, O]) PostorderTraversal(fn func(N)) {
	PostorderTraversal(t.Roots, *new(O), fn)
}

func (t *Trees[ID, T, N, O]) LevelOrderTraversal(fn func(N)) {
	LevelOrderTraversal(t.Roots, *new(O), fn)
}

func (t *Trees[ID, T, N, O]) ForEach(fn func(N)) {
	t.PreorderTraversal(fn)
}

// PreorderTrees wraps Trees to provide pre-order traversal of node values.
// This type implements Container[T] instead of Container[N],
// allowing direct access to stored values rather than nodes.
type PreorderTrees[ID comparable, T any, N TreeNode[ID, T], O NodeOperator[ID, T, N]] struct {
	Trees[ID, T, N, O]
}

// ForEach executes the given function for each node's value in pre-order traversal.
// Implements Container[T].ForEach().
func (t *PreorderTrees[ID, T, N, O]) ForEach(fn func(T)) {
	t.PreorderTraversal(func(node N) {
		fn(node.GetValue())
	})
}

// PostorderTrees wraps Trees to provide post-order traversal of node values.
// This type implements Container[T] instead of Container[N],
// allowing direct access to stored values rather than nodes.
type PostorderTrees[ID comparable, T any, N TreeNode[ID, T], O NodeOperator[ID, T, N]] struct {
	Trees[ID, T, N, O]
}

// ForEach executes the given function for each node's value in post-order traversal.
// Implements Container[T].ForEach().
func (t *PostorderTrees[ID, T, N, O]) ForEach(fn func(T)) {
	t.PostorderTraversal(func(node N) {
		fn(node.GetValue())
	})
}

// LevelOrderTrees wraps Trees to provide level-order traversal of node values.
// This type implements Container[T] instead of Container[N],
// allowing direct access to stored values rather than nodes.
type LevelOrderTrees[ID comparable, T any, N TreeNode[ID, T], O NodeOperator[ID, T, N]] struct {
	Trees[ID, T, N, O]
}

// ForEach executes the given function for each node's value in level-order traversal.
// Implements Container[T].ForEach().
func (t *LevelOrderTrees[ID, T, N, O]) ForEach(fn func(T)) {
	t.LevelOrderTraversal(func(node N) {
		fn(node.GetValue())
	})
}

var _ icontainer.Container[*OrderedNode[int, string]] = (*Trees[int, string, *OrderedNode[int, string], OrderedNodeOperator[int, string]])(nil)
var _ icontainer.Container[string] = (*PreorderTrees[int, string, *OrderedNode[int, string], OrderedNodeOperator[int, string]])(nil)
var _ icontainer.Container[string] = (*PostorderTrees[int, string, *OrderedNode[int, string], OrderedNodeOperator[int, string]])(nil)
var _ icontainer.Container[string] = (*LevelOrderTrees[int, string, *OrderedNode[int, string], OrderedNodeOperator[int, string]])(nil)
