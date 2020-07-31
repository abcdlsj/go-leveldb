// Package skiplist is the skiplist implementation by me
package skiplist

const (
	MaxHeight = 12
	Branching = 4
)

type SkipList struct {
	maxHeight int
	head      *Node
}
