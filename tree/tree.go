package tree

type ITree interface {
	Len() int
	Get(key int) Value
	Put(key int, val Value)
	Delete(key int)
	Traverse() []Value
	Deepth() int
}

type Tree struct {
	Root *Node
	Size int
}

func New() *Tree {
	return &Tree{}
}

func (t *Tree) Len() int {
	return t.Size
}

func (t *Tree) Get(key int) Value {
	node := t.Root.get(key)
	if node != nil {
		return node.Val
	}
	return nil
}

func (t *Tree) Put(key int, val Value) {
	if t.Root == nil {
		t.Root = newNode(key, val, t, nil)
		t.Root.Color = Black
		t.Size++
	} else {
		t.Root.put(key, val)
	}
}

func (t *Tree) Delete(key int) {
	t.Root.delete(key)
}

func (t *Tree) Traverse() []Value {
	nodes := t.Root.traverse()
	var res []Value
	for _, v := range nodes {
		res = append(res, v.Val)
	}
	return res
}

func (t *Tree) Deepth() int {
	return t.Root.deepth()
}
