package bst

type BinaryTreeNode struct {
	value int
	left *BinaryTreeNode
	right *BinaryTreeNode
}

type BST struct {
	root *BinaryTreeNode
}

// instantiate an empty BST
func New() *BST {
	return &BST{}
}

// insert a new element to BST
func (b *BST) Insert(value int)  {
	ptr := b.root
	node := &BinaryTreeNode{value: value}
	if ptr==nil {
		b.root = node
		return
	}
	var prev *BinaryTreeNode = nil
	for ptr!=nil {
		prev = ptr
		if value > ptr.value {
			ptr = ptr.right
		}else {
			ptr = ptr.left
		}
	}
	if value > prev.value {
		prev.right = node
	}else {
		prev.left = node
	}
}

// check if element in BST
func (b *BST) Contains(value int) bool {
	ptr := b.root
	for ptr!=nil {
		if value==ptr.value {
			return true
		} else if value > ptr.value {
			ptr = ptr.right
		}else {
			ptr = ptr.left
		}
	}
	return false
}

func inorderHelper(root *BinaryTreeNode, result *[]int) {
	if root!=nil {
		inorderHelper(root.left, result)
		(*result) = append((*result), root.value)
		inorderHelper(root.right, result)
	}
}

// inorder traversal for BST
func (b *BST) Inorder() []int {
	var inorder []int 
	inorderHelper(b.root, &inorder)
	return inorder
}

func preorderHelper(root *BinaryTreeNode, result *[]int) {
	if root!=nil {
		(*result) = append((*result), root.value)
		preorderHelper(root.left, result)
		preorderHelper(root.right, result)
	}
}

// preorder traversal for BST
func (b *BST) Preorder() []int {
	var preorder []int 
	preorderHelper(b.root, &preorder)
	return preorder
}

func postorderHelper(root *BinaryTreeNode, result *[]int) {
	if root!=nil {
		postorderHelper(root.left, result)
		postorderHelper(root.right, result)
		(*result) = append((*result), root.value)
	}
}

// postorder traversal for BST
func (b *BST) Postorder() []int {
	var postorder []int 
	postorderHelper(b.root, &postorder)
	return postorder
}
