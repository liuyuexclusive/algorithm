package tree

import "math"

const (
	Red   Color = true
	Black Color = false
)

type Color bool

type Value interface{}

type Node struct {
	Tree                *Tree
	Key                 int
	Val                 Value
	Left, Right, Parent *Node
	Color               Color
}

func newNode(key int, val Value, tree *Tree, parent *Node) *Node {
	res := &Node{Key: key, Val: val, Tree: tree, Parent: parent}
	res.Color = Red
	return res
}

func (n *Node) get(key int) *Node {
	if n == nil {
		return nil
	}
	switch {
	case key < n.Key:
		return n.Left.get(key)
	case key > n.Key:
		return n.Right.get(key)
	case key == n.Key:
		return n
	}
	return nil
}

func (n *Node) put(key int, val Value) *Node {
	if n == nil {
		return nil
	}

	nn := newNode(key, val, n.Tree, n)
	if key == 3 {

	}

	switch {
	case key == n.Key:
		n.Val = val
		nn = n
	case key < n.Key:
		if n.Left == nil {
			nn.Tree.Size++
			n.Left = nn
			n.rotate()
			break
		}
		return n.Left.put(key, val)
	case key > n.Key:
		if n.Right == nil {
			nn.Tree.Size++
			n.Right = nn
			n.rotate()
			break
		}
		return n.Right.put(key, val)
	}

	return nn
}

func (n *Node) rotate() {
	if !n.Left.isRed() && n.Right.isRed() {
		n.rotateLeft().rotate()
	} else if n.Parent.isRed() && n.Left.isRed() {
		n.Parent.rotateRight().rotate()
	} else if n.Left.isRed() && n.Right.isRed() {
		n.reverseColor().rotate()
	}
}

func (n *Node) delete(key int) *Node {
	if n == nil {
		return nil
	}
	switch {
	case n.Key == key:
		n.Tree.Size--
		if n.Parent == nil {
			if n.Left == nil && n.Right == nil {
				n.Tree.Root = nil
			} else if n.Left != nil && n.Right == nil {
				n.Left.Parent = nil
				n.Left.Color = Black
				n.Tree.Root = n.Left
			} else {
				next := n.next()
				next.deleteFromParent(next.Right)
				next.Left = n.Left
				next.Right = n.Right
				next.Parent = nil
				next.Color = Black
				n.Tree.Root = next
			}
		} else {
			if n.Left == nil && n.Right == nil {
				n.deleteFromParent(nil)
			} else if n.Left != nil && n.Right == nil {
				n.deleteFromParent(n.Left)
				n.Left.Parent = n.Parent
			} else {
				next := n.next()
				next.deleteFromParent(next.Right)
				next.Left = n.Left
				next.Right = n.Right
				next.Parent = n.Parent
				n.deleteFromParent(next)
			}
		}
		return n
	case key < n.Key:
		return n.Left.delete(key)
	case key > n.Key:
		return n.Right.delete(key)
	}
	return nil
}

func (n *Node) deleteFromParent(replacement *Node) {
	if n == nil || n.Parent == nil {
		return
	}
	if n.Parent.Left != nil && n.Key == n.Parent.Left.Key {
		n.Parent.Left = replacement
	} else {
		n.Parent.Right = replacement
	}
}

// find next Node
func (n *Node) next() *Node {
	if n == nil || n.Right == nil {
		return nil
	}
	n = n.Right
	if n.Left != nil {
		return n.Left.next()
	}
	return n
}

func (n *Node) traverse() []*Node {
	res := make([]*Node, 0)
	if n == nil {
		return res
	}
	res = append(res, n.Left.traverse()...)
	res = append(res, n)
	res = append(res, n.Right.traverse()...)
	return res
}

func (n *Node) rotateLeft() *Node {
	r := n.Right
	n.Right = r.Left
	r.Left = n
	if n.Parent == nil {
		r.Parent = nil
		r.Tree.Root = r
	} else {
		r.Parent = n.Parent
		n.deleteFromParent(r)
	}
	n.Parent = r
	r.Color = n.Color
	n.Color = Red
	return r
}

func (n *Node) rotateRight() *Node {
	l := n.Left
	n.Left = l.Right
	l.Right = n
	if n.Parent == nil {
		l.Parent = nil
		l.Tree.Root = l
	} else {
		l.Parent = n.Parent
		n.deleteFromParent(l)
	}
	l.Color = n.Color
	n.Color = Red
	return l
}

func (n *Node) reverseColor() *Node {
	n.Left.Color = Black
	n.Right.Color = Black
	if n.Parent == nil {
		n.Color = Black
	} else {
		n.Color = Red
		if n.Parent != nil {
			n = n.Parent
		}
	}
	return n
}

func (n *Node) isRed() bool {
	if n == nil {
		return false
	}
	return n.Color == Red
}

func (n *Node) deepth() int {
	if n == nil {
		return 0
	}
	return int(math.Max(float64(n.Left.deepth()), float64(n.Right.deepth()))) + 1
}
