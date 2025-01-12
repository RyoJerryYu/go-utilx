package treex

import "github.com/RyoJerryYu/go-utilx/pkg/container/icontainer"

type TreeNode[ID comparable, T any] interface {
	GetId() ID
	GetPath() []ID
	GetValue() T
	Size() int
}

type NodeOperator[ID comparable, T any, N TreeNode[ID, T]] interface {
	Parent(node N) N
	Children(node N) []N
	AddChild(node N, child N)
}

func PreorderTraversal[ID comparable, T any, N TreeNode[ID, T], O NodeOperator[ID, T, N]](roots []N, operator O, fn func(N)) {
	for _, root := range roots {
		fn(root)
		PreorderTraversal(operator.Children(root), operator, fn)
	}
}

func PostorderTraversal[ID comparable, T any, N TreeNode[ID, T], O NodeOperator[ID, T, N]](roots []N, operator O, fn func(N)) {
	for _, root := range roots {
		PostorderTraversal(operator.Children(root), operator, fn)
		fn(root)
	}
}

func LevelOrderTraversal[ID comparable, T any, N TreeNode[ID, T], O NodeOperator[ID, T, N]](roots []N, operator O, fn func(N)) {
	queue := roots
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		fn(node)
		queue = append(queue, operator.Children(node)...)
	}
}

type Trees[ID comparable, T any, N TreeNode[ID, T], O NodeOperator[ID, T, N]] struct {
	Roots []N
}

func New[ID comparable, T any, N TreeNode[ID, T], O NodeOperator[ID, T, N]]() *Trees[ID, T, N, O] {
	return &Trees[ID, T, N, O]{}
}
func (t *Trees[ID, T, N, O]) GetRoots() []N { return t.Roots }

func (t *Trees[ID, T, N, O]) Len() int {
	cnt := 0
	for _, root := range t.Roots {
		cnt += root.Size()
	}
	return cnt
}

func (t *Trees[ID, T, N, O]) IsEmpty() bool {
	return len(t.Roots) == 0
}

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

type PreorderTrees[ID comparable, T any, N TreeNode[ID, T], O NodeOperator[ID, T, N]] struct {
	Trees[ID, T, N, O]
}

func (t *PreorderTrees[ID, T, N, O]) ForEach(fn func(T)) {
	t.PreorderTraversal(func(node N) {
		fn(node.GetValue())
	})
}

type PostorderTrees[ID comparable, T any, N TreeNode[ID, T], O NodeOperator[ID, T, N]] struct {
	Trees[ID, T, N, O]
}

func (t *PostorderTrees[ID, T, N, O]) ForEach(fn func(T)) {
	t.PostorderTraversal(func(node N) {
		fn(node.GetValue())
	})
}

type LevelOrderTrees[ID comparable, T any, N TreeNode[ID, T], O NodeOperator[ID, T, N]] struct {
	Trees[ID, T, N, O]
}

func (t *LevelOrderTrees[ID, T, N, O]) ForEach(fn func(T)) {
	t.LevelOrderTraversal(func(node N) {
		fn(node.GetValue())
	})
}

var _ icontainer.Container[*OrderedNode[int, string]] = (*Trees[int, string, *OrderedNode[int, string], OrderedNodeOperator[int, string]])(nil)
var _ icontainer.Container[string] = (*PreorderTrees[int, string, *OrderedNode[int, string], OrderedNodeOperator[int, string]])(nil)
var _ icontainer.Container[string] = (*PostorderTrees[int, string, *OrderedNode[int, string], OrderedNodeOperator[int, string]])(nil)
var _ icontainer.Container[string] = (*LevelOrderTrees[int, string, *OrderedNode[int, string], OrderedNodeOperator[int, string]])(nil)
