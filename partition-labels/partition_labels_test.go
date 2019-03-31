package partition_labels

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func testCase(t *testing.T, f func (string)[]int) {
	assert.Equal(t, []int{9,7,8}, f("ababcbacadefegdehijhklij"))
}

func Test(t *testing.T) {
	testCase(t, partitionLabels)
}

