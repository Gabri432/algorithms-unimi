package algorithmsunimi

import "testing"

func TestPrintTreeSummary(t *testing.T) {
	tr := &tree{nil}
	tr.root = &treeNode{nil, nil, 78}
	tr.root.left = newTreeNode(54)
	tr.root.right = newTreeNode(21)
	tr.root.left.right = newTreeNode(90)
	tr.root.left.right.left = newTreeNode(19)
	tr.root.left.right.right = newTreeNode(95)
	tr.root.right.left = newTreeNode(16)
	tr.root.right.left.left = newTreeNode(5)
	tr.root.right.right = newTreeNode(19)
	tr.root.right.right.left = newTreeNode(56)
	tr.root.right.right.right = newTreeNode(43)
	printTreeSummary(tr.root, 0)
}
