package redblack

type Color int

const (
	Black Color = iota
	Red
)

type Node[T comparable] struct {
	Left, Right, Parent *Node[T]
	Key                 T
	Color               Color
}

type Tree[T comparable] struct {
	Root Node[T]
}

func (t Tree[T]) Search(key T) Node[T] {
}
