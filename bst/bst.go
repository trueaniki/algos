package bst

import (
	"math"
	"sort"
)

type BST struct {
	Left  *BST
	Right *BST
	Key   int
	Value interface{}
	// LessFun func(int, int) bool
}

func search(n *BST, key int) interface{} {
	if n.Key == key {
		return n.Value
	}
	if n.Key < key {
		return search(n.Left, key)
	}
	return search(n.Right, key)
}

func (t *BST) Search(key int) interface{} {
	return search(t, key)
}

func (t *BST) InOrder(cb func(n *BST)) {
	inOrder(t, cb)
}

func inOrder(n *BST, cb func(n *BST)) {
	if n == nil {
		return
	}
	inOrder(n.Left, cb)
	cb(n)
	inOrder(n.Right, cb)
}

func (t *BST) PreOrder(cb func(n *BST)) {
	preOrder(t, cb)
}

func preOrder(n *BST, cb func(n *BST)) {
	if n == nil {
		return
	}
	cb(n)
	preOrder(n.Left, cb)
	preOrder(n.Right, cb)
}

func (t *BST) PostOrder(cb func(n *BST)) {
	postOrder(t, cb)
}

func postOrder(n *BST, cb func(n *BST)) {
	if n == nil {
		return
	}
	postOrder(n.Left, cb)
	postOrder(n.Right, cb)
	cb(n)
}

func (t *BST) AllKeys() []int {
	return []int{}
}

func (t *BST) AVL() *BST {
	nodes := []struct {
		Key   int
		Value interface{}
	}{}

	t.InOrder(func(n *BST) {
		nodes = append(nodes, struct {
			Key   int
			Value interface{}
		}{
			n.Key,
			n.Value,
		})
	})
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].Key < nodes[j].Key
	})

	avl := &BST{}
	addInBalance(avl, 0, len(nodes)-1, nodes)
	return avl
}

func addInBalance(n *BST, start, end int, nodes []struct {
	Key   int
	Value interface{}
}) {
	if end-start == 1 {
		return
	}

	median := int(math.Floor((float64(end)-float64(start))/2)) + start
	// median := (end - start) / 2
	// fmt.Printf("%+v", struct {
	// 	Start  int
	// 	End    int
	// 	Median int
	// }{
	// 	Start:  start,
	// 	End:    end,
	// 	Median: median,
	// })
	n.Key = nodes[median].Key
	n.Value = nodes[median].Value

	addInBalance(n.Left, start, median, nodes)
	addInBalance(n.Right, median, end, nodes)
}

func findForAdding(n *BST, key int) *BST {
	if n.Key == key {
		panic("Key exists!!!")
	}

	if key < n.Key {
		if n.Left == nil {
			n.Left = &BST{}
			return n.Left
		}
		return findForAdding(n.Left, key)
	} else {
		if n.Right == nil {
			n.Right = &BST{}
			return n.Right
		}
		return findForAdding(n.Right, key)
	}
}

func (t *BST) Add(key int, value interface{}) {
	var node *BST

	if t.Key == 0 && t.Value == nil && t.Right == nil && t.Left == nil {
		node = t
	} else {
		node = findForAdding(t, key)
	}
	node.Left = nil
	node.Right = nil
	node.Key = key
	node.Value = value
}
