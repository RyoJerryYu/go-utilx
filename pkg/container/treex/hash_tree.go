package treex

import "github.com/RyoJerryYu/go-utilx/pkg/container/slicex"

// HashNode represents a node in a tree structure that uses a hash map to store children.
// This implementation allows for O(1) child lookup by ID.
// Type parameters:
//   - ID: The type of node identifier, must be comparable
//   - T: The type of value stored in the node
type HashNode[ID comparable, T any] struct {
	Id       ID
	Path     []ID // Path from root to this node, including this node.
	Parent   *HashNode[ID, T]
	Children map[ID]*HashNode[ID, T]
	Value    T
}

func (n *HashNode[ID, T]) IsRoot() bool                         { return n.Parent == nil }
func (n *HashNode[ID, T]) IsLeaf() bool                         { return len(n.Children) == 0 }
func (n *HashNode[ID, T]) GetId() ID                            { return n.Id }
func (n *HashNode[ID, T]) GetPath() []ID                        { return n.Path }
func (n *HashNode[ID, T]) GetParent() *HashNode[ID, T]          { return n.Parent }
func (n *HashNode[ID, T]) GetChildren() map[ID]*HashNode[ID, T] { return n.Children }
func (n *HashNode[ID, T]) GetValue() T                          { return n.Value }
func (n *HashNode[ID, T]) Size() int {
	cnt := 1
	for _, child := range n.Children {
		cnt += child.Size()
	}
	return cnt
}

// HashNodeOperator provides operations for HashNode trees.
// It implements the NodeOperator interface for HashNode types.
type HashNodeOperator[ID comparable, T any] struct{}

func (HashNodeOperator[ID, T]) Parent(node *HashNode[ID, T]) *HashNode[ID, T] {
	return node.Parent
}

func (HashNodeOperator[ID, T]) Children(node *HashNode[ID, T]) []*HashNode[ID, T] {
	return slicex.FromValue(node.Children)
}

func (HashNodeOperator[ID, T]) AddChild(node *HashNode[ID, T], child *HashNode[ID, T]) {
	child.Parent = node
	child.Path = append(node.Path, child.Id)
	node.Children[child.Id] = child
}

var _ NodeOperator[any, any, *HashNode[any, any]] = HashNodeOperator[any, any]{}
