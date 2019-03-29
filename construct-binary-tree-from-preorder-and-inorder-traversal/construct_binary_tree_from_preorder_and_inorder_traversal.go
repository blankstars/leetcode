package construct_binary_tree_from_preorder_and_inorder_traversal

/**
	https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/

	题目:
	根据一棵树的前序遍历与中序遍历构造二叉树。
	注意:
	你可以假设树中没有重复的元素。

	例如，给出
	前序遍历 preorder = [3,9,20,15,7]
	中序遍历 inorder = [9,3,15,20,7]
	返回如下的二叉树：
        3
       / \
      9  20
        /  \
       15   7
 */

/**
* Definition for a binary tree node.
*/
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if preorder == nil || len(preorder) == 0 {
		return nil
	}

	root := preorder[0]
	var leftTreeNode, rightTreeNode *TreeNode
	for i, node := range inorder {
		var leftPreOrder, rightPreOrder []int
		var leftInOrder, rightInOrder []int

		if node != root {
			continue
		}

		if i == 0 {
			leftPreOrder = nil
			leftInOrder = nil
		} else {
			leftPreOrder = preorder[1:i+1]
			leftInOrder = inorder[:i]
		}
		if i == len(inorder) {
			rightPreOrder = nil
			rightInOrder = nil
		} else {
			rightPreOrder = preorder[i+1:]
			rightInOrder = inorder[i+1:]
		}

		leftTreeNode = buildTree(leftPreOrder, leftInOrder)
		rightTreeNode = buildTree(rightPreOrder, rightInOrder)
	}

	return &TreeNode{Val:root, Left:leftTreeNode, Right:rightTreeNode}
}