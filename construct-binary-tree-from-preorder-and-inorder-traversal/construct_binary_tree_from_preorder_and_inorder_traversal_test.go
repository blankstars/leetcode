package construct_binary_tree_from_preorder_and_inorder_traversal

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"gopkg.in/eapache/queue.v1"
	"testing"
)

func conv(q *queue.Queue) []string {
	array := make([]string, 0)

	for {
		if q.Length() == 0 {
			break
		}

		node := q.Remove().(*TreeNode)
		if node == nil {
			array = append(array, "null")
		} else {
			array = append(array, fmt.Sprintf("%v", node.Val))
			q.Add(node.Left)
			q.Add(node.Right)
		}
	}

	return array
}


func testCase(t *testing.T, f func ([]int, []int) *TreeNode) {
	q := queue.New()
	q.Add(f([]int{3,9,20,15,7}, []int{9,3,15,20,7}))
	assert.Equal(t, []string{"3","9","20","null","null","15","7","null","null","null","null"}, conv(q))
}

func Test(t *testing.T) {
	testCase(t, buildTree)
}
