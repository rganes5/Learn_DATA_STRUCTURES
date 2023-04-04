package main

import (
	"fmt"
	"math"
)

//Tree Structure

type tree struct {
	root *treeNode
}

type treeNode struct {
	data  int
	left  *treeNode
	right *treeNode
}

//insertion

func (t *tree) insert(value int) {
	if t.root == nil {
		t.root = &treeNode{data: value}
	} else {
		t.root.Insert(value)
	}
}
func (n *treeNode) Insert(value int) {
	if value < n.data {
		if n.left == nil {
			n.left = &treeNode{data: value}
		} else {
			n.left.Insert(value)
		}
	} else if value > n.data {
		if n.right == nil {
			n.right = &treeNode{data: value}
		} else {
			n.right.Insert(value)
		}
	}
}

//contains

func (t *tree) contains(value int) bool {
	if t.root == nil {
		return false
	} else {
		return t.root.containsNode(value)
	}
}
func (n *treeNode) containsNode(value int) bool {
	if n == nil {
		return false
	} else if value < n.data {
		return n.left.containsNode(value)
	} else if value > n.data {
		return n.right.containsNode(value)
	} else {
		return true
	}
}

//Deletion

func (t *tree) delete(value int) {
	if t.root == nil {
		fmt.Println("Tree empty")
	} else {
		if t.root.containsNode(value) {
			t.root.deleteHelper(value)
		}
	}
}
func (n *treeNode) deleteHelper(value int) *treeNode {
	if value < n.data {
		n.left = n.left.deleteHelper(value)
	} else if value > n.data {
		n.right = n.right.deleteHelper(value)
	} else {
		if n.left == nil && n.right == nil {
			return nil
		} else if n.left == nil {
			return n.right
		} else if n.right == nil {
			return n.left
		} else {
			//We are finding the minimum value from the RHS
			tempNode := findSuccessor(n.right)
			n.data = tempNode.data
			n.right = n.right.deleteHelper(tempNode.data)
		}
	}
	return n
}

//Helper to find the next successor to be replaced by the replaced node in tree

func findSuccessor(tempNode *treeNode) *treeNode {
	for tempNode.left != nil {
		tempNode = tempNode.left
	}
	return tempNode
}

//Tree traversal methods

//In-order- left,root,right

func inOrder(root *treeNode) {
	if root == nil {
		return
	} else {
		inOrder(root.left)
		fmt.Print(root.data, " ")
		inOrder(root.right)
	}
}

// Post-Order - left,right,root
func postOrder(root *treeNode) {
	if root == nil {
		return
	} else {
		postOrder(root.left)
		postOrder(root.right)
		fmt.Print(root.data, " ")

	}
}

// Pre-Order - root,left,right
func preOrder(root *treeNode) {
	if root == nil {
		return
	} else {
		fmt.Print(root.data, " ")
		preOrder(root.left)
		preOrder(root.right)
	}
}

// Find the closest value to a given number in a Tree.

func (t *tree) closestValue(value int) {
	if t.root == nil {
		fmt.Println("Tree is empty")
	} else {
		fmt.Println(t.root.helperClosestValue(value))
	}
}
func (n *treeNode) helperClosestValue(value int) int {
	difference := math.MaxInt
	var currentDifference, currentData int
	for n != nil {
		if n.data < value {
			currentDifference = value - n.data
			if currentDifference < difference {
				difference = currentDifference
				currentData = n.data
			}
			n = n.right
		} else if n.data > value {
			currentDifference = n.data - value
			if currentDifference < difference {
				difference = currentDifference
				currentData = n.data
			}
			n = n.left
		} else {
			return n.data
		}
	}
	return currentData
}

// Validate whether it is a binary tree or not
func (t *tree) Validate() bool {
	min := math.MinInt
	max := math.MaxInt
	return t.root.helperValidate(min, max)
}

func (n *treeNode) helperValidate(min, max int) bool {
	if n == nil {
		return true
	} else if n.data < min || n.data > max {
		return false
	}
	return n.left.helperValidate(min, n.data-1) && n.right.helperValidate(n.data+1, max)
}

// MAIN FUNCTION
func main() {
	tree := tree{}
	array := []int{5, 4, 3, 10, 8, 6, 1}
	for _, v := range array {
		tree.insert(v)
	}
	fmt.Println(tree.contains(10))
	// tree.delete(10)
	// fmt.Println(tree.contains(10))
	fmt.Println("INORDER")
	inOrder(tree.root)
	fmt.Println("")
	fmt.Println("POSTORDER")
	postOrder(tree.root)
	fmt.Println("")
	fmt.Println("PREORDER")
	preOrder(tree.root)
	fmt.Println("")
	tree.closestValue(9)
	fmt.Println(tree.Validate())

}
