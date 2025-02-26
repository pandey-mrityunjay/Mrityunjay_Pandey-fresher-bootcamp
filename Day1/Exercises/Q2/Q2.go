package main

import (
	"fmt"
)

// TreeNode struct
type TreeNode struct {
	Value string
	Left  *TreeNode
	Right *TreeNode
}

// Preorder Traversal (Root -> Left -> Right)
func Preorder(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Print(node.Value, " ")
	Preorder(node.Left)
	Preorder(node.Right)
}

// Postorder Traversal (Left -> Right -> Root)
func Postorder(node *TreeNode) {
	if node == nil {
		return
	}
	Postorder(node.Left)
	Postorder(node.Right)
	fmt.Print(node.Value, " ")
}

// Main function
func main() {
	// Manually constructing tree for a + (b - c)
	root := &TreeNode{Value: "+"}
	root.Left = &TreeNode{Value: "a"}
	root.Right = &TreeNode{Value: "-"}
	root.Right.Left = &TreeNode{Value: "b"}
	root.Right.Right = &TreeNode{Value: "c"}
	fmt.Println("Preorder Traversal:")
	Preorder(root)
	fmt.Println("\nPostorder Traversal:")
	Postorder(root)
}
