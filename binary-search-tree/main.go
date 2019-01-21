package main

import (
	"fmt"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type BinarySearchTree struct {
	Root *Node
}

func (bst *BinarySearchTree) Insert(value int) {
	node := &Node{Value: value}

	if bst.Root == nil {
		bst.Root = node
		return
	}

	current := bst.Root

	for {
		// 若已有相同值則不處理。
		if node.Value == current.Value {
			return
		}

		if node.Value < current.Value {
			if current.Left == nil {
				current.Left = node
				return
			}
			current = current.Left
		} else {
			if current.Right == nil {
				current.Right = node
				return
			}
			current = current.Right
		}
	}
}

func (bst *BinarySearchTree) BreadthFirstSearch() []int {
	node := bst.Root
	result := make([]int, 0)
	queue := make([]*Node, 0)
	queue = append(queue, node)

	for len(queue) > 0 {
		node = queue[0]
		queue = queue[1:]
		result = append(result, node.Value)

		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return result
}

func (bst *BinarySearchTree) DepthFirstSearchPreOrder() []int {
	result := make([]int, 0)
	var traverse func(node *Node)
	traverse = func(node *Node) {
		result = append(result, node.Value)

		if node.Left != nil {
			traverse(node.Left)
		}

		if node.Right != nil {
			traverse(node.Right)
		}
	}

	traverse(bst.Root)
	return result
}

func (bst *BinarySearchTree) DepthFirstSearchPostOrder() []int {
	result := make([]int, 0)
	var traverse func(node *Node)
	traverse = func(node *Node) {
		if node.Left != nil {
			traverse(node.Left)
		}

		if node.Right != nil {
			traverse(node.Right)
		}

		result = append(result, node.Value)
	}

	traverse(bst.Root)
	return result
}

func (bst *BinarySearchTree) DepthFirstSearchInOrder() []int {
	result := make([]int, 0)
	var traverse func(node *Node)
	traverse = func(node *Node) {
		if node.Left != nil {
			traverse(node.Left)
		}

		result = append(result, node.Value)

		if node.Right != nil {
			traverse(node.Right)
		}
	}

	traverse(bst.Root)
	return result
}

func main() {
	fmt.Println("Binary Search Tree")
	bst := BinarySearchTree{}
	bst.Insert(10)
	bst.Insert(6)
	bst.Insert(15)
	bst.Insert(3)
	bst.Insert(8)
	bst.Insert(20)
	fmt.Println("Breadth First Search", bst.BreadthFirstSearch())
	fmt.Println("Depth First Search Pre-Order", bst.DepthFirstSearchPreOrder())
	fmt.Println("Depth First Search Post-Order", bst.DepthFirstSearchPostOrder())
	fmt.Println("Depth First Search In-Order", bst.DepthFirstSearchInOrder())
}
