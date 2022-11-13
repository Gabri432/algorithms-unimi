package algorithmsunimi

import (
	"fmt"
	"math/rand"
	"time"
)

type treeNode struct {
	left  *treeNode
	right *treeNode
	item  int
}

type tree struct {
	root *treeNode
}

func newTreeNode(val int) *treeNode {
	return &treeNode{nil, nil, val}
}

func generateRamdom(length int) []int {

	mySlice := make([]int, length)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		mySlice = append(mySlice, rand.Intn(100))
	}
	return mySlice
}

func treeInOrder(t *tree) {
	inorder(t.root)
	fmt.Println()
}

func treePreOrder(t *tree) {
	preorder(t.root)
	fmt.Println()
}

func treePostOrder(t *tree) {
	postorder(t.root)
	fmt.Println()
}

func inorder(node *treeNode) {
	if node == nil {
		return
	}
	inorder(node.left)
	fmt.Print(node.item)
	inorder(node.right)
}

func preorder(node *treeNode) {
	if node == nil {
		return
	}
	fmt.Print(node.item)
	preorder(node.left)
	preorder(node.right)
}

func postorder(node *treeNode) {
	if node == nil {
		return
	}
	postorder(node.left)
	postorder(node.right)
	fmt.Print(node.item)
}

func stampaAlberoASommario(node *treeNode, spaces int) {
	for i := 0; i < spaces; i++ {
		fmt.Print(" ")
	}
	fmt.Print("*")
	fmt.Println(node.item)
	if node.left != nil || node.right != nil {
		stampaAlberoASommario(node.left, spaces+1)
		stampaAlberoASommario(node.right, spaces+1)
	}
}
