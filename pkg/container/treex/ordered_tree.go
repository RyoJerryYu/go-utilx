package treex

// OrderedNode represents a node in a tree structure that maintains children in an ordered slice.
// This implementation preserves the order of child nodes as they are added.
// Type parameters:
//   - ID: The type of node identifier, must be comparable
//   - T: The type of value stored in the node
type OrderedNode[ID comparable, T any] struct {
	Id       ID
	Path     []ID // Path from root to this node, including this node.
	Parent   *OrderedNode[ID, T]
	Children []*OrderedNode[ID, T]
	Value    T
}

func (n *OrderedNode[ID, T]) IsRoot() bool                       { return n.Parent == nil }
func (n *OrderedNode[ID, T]) IsLeaf() bool                       { return len(n.Children) == 0 }
func (n *OrderedNode[ID, T]) GetId() ID                          { return n.Id }
func (n *OrderedNode[ID, T]) GetPath() []ID                      { return n.Path }
func (n *OrderedNode[ID, T]) GetParent() *OrderedNode[ID, T]     { return n.Parent }
func (n *OrderedNode[ID, T]) GetChildren() []*OrderedNode[ID, T] { return n.Children }
func (n *OrderedNode[ID, T]) GetValue() T                        { return n.Value }
func (n *OrderedNode[ID, T]) Size() int {
	cnt := 1
	for _, child := range n.Children {
		cnt += child.Size()
	}
	return cnt
}

// OrderedNodeOperator provides operations for OrderedNode trees.
// It implements the NodeOperator interface for OrderedNode types.
type OrderedNodeOperator[ID comparable, T any] struct{}

func (OrderedNodeOperator[ID, T]) Parent(node *OrderedNode[ID, T]) *OrderedNode[ID, T] {
	return node.Parent
}
func (OrderedNodeOperator[ID, T]) Children(node *OrderedNode[ID, T]) []*OrderedNode[ID, T] {
	return node.Children
}

func (OrderedNodeOperator[ID, T]) AddChild(node *OrderedNode[ID, T], child *OrderedNode[ID, T]) {
	child.Parent = node
	child.Path = append(node.Path, child.Id)
	node.Children = append(node.Children, child)
}

var _ NodeOperator[any, any, *OrderedNode[any, any]] = OrderedNodeOperator[any, any]{}
