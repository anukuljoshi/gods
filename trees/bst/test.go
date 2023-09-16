package bst

import "fmt"

func BSTTest() {
	bst := New()
	value := 5
	fmt.Println("*********** BST Insert", value)
	bst.Insert(value)
	value = 6
	fmt.Println("*********** BST Insert", value)
	bst.Insert(value)
	value = 8
	fmt.Println("*********** BST Insert", value)
	bst.Insert(value)
	value = 2
	fmt.Println("*********** BST Insert", value)
	bst.Insert(value)
	value = 0
	fmt.Println("*********** BST Insert", value)
	bst.Insert(value)
	value = 1
	fmt.Println("*********** BST Insert", value)
	bst.Insert(value)
	value = 4
	fmt.Println("*********** BST Insert", value)
	bst.Insert(value)
	fmt.Println("*********** BST Inorder", bst.Inorder())
	fmt.Println("*********** BST Preorder", bst.Preorder())
	fmt.Println("*********** BST Postorder", bst.Postorder())
	fmt.Println("*********** BST Contains", bst.Contains(4))
	fmt.Println("*********** BST Contains", bst.Contains(1))
	fmt.Println("*********** BST Contains", bst.Contains(100))
}